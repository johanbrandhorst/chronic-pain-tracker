package api

import (
	"context"
	"database/sql"
	"net"
	"net/url"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/cenk/backoff"
	"github.com/elgris/sqrl"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/logrusadapter"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"

	"github.com/johanbrandhorst/chronic-pain-tracker/api/internal"
	"github.com/johanbrandhorst/chronic-pain-tracker/proto"
)

//go:generate protoc -I../proto -I../vendor/github.com/googleapis/googleapis/ --go_out=plugins=grpc:$GOPATH/src --grpc-gateway_out=logtostderr=true:../proto/ ../proto/api.proto

var _ proto.PainTrackerServer = (*Server)(nil)
var _ proto.MonitorServer = (*Server)(nil)

type Server struct {
	db *sql.DB
	sb sqrl.StatementBuilderType
}

func NewServer(logger *logrus.Logger, pqURL url.URL) (*Server, error) {
	bk := backoff.NewExponentialBackOff()
	bk.MaxElapsedTime = 0 // Ensure we never stop
	// Retry dialling until connection established
	dialFunc := func(network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			KeepAlive: 5 * time.Minute,
			Timeout:   5 * time.Second,
		}
		var conn net.Conn
		connFn := func() error {
			var err error
			conn, err = dialer.Dial(network, addr)
			return err
		}
		// Retry in perpetuity
		_ = backoff.RetryNotify(
			connFn,
			bk,
			func(err error, next time.Duration) {
				logger.WithError(err).Warnf("Failed to connect to postgres, retrying in %v", next.Truncate(time.Millisecond))
			},
		)

		return conn, nil
	}

	driverConfig := stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			Logger: logrusadapter.NewLogger(logger),
			Dial:   dialFunc,
		},
	}
	stdlib.RegisterDriverConfig(&driverConfig)
	db, err := sql.Open("pgx", driverConfig.ConnectionString(pqURL.String()))
	if err != nil {
		return nil, err
	}

	sb := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)

	// Retry ensureSchema separately, as it may error with a temporary error
	// even after the dialing has completed.
	err = backoff.RetryNotify(
		func() error {
			// Ensures schema exists
			err := ensureSchema(db, sb) // nolint: vetshadow
			if err != nil {
				e, ok := err.(pgx.PgError)
				if ok && e.Code == psqlCannotConnectNow {
					// Retry connection on this one specific error
					return e
				}

				// Other errors are considered permanent
				return backoff.Permanent(err)
			}

			return nil
		},
		bk,
		func(err error, next time.Duration) {
			logger.WithError(err).Warnf("Failed to call postgres client, retrying in %v", next.Truncate(time.Millisecond))
		},
	)
	if err != nil {
		return nil, err
	}

	return &Server{
		db: db,
		sb: sb,
	}, nil
}

func (s *Server) SetPainLevel(ctx context.Context, req *proto.PainUpdate) (*empty.Empty, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var stressed bool
	var flare bool
	q := s.sb.Select(
		internal.StressedColumn,
		internal.FlareColumn,
	).From(
		internal.EventsTable,
	).OrderBy(
		internal.TimestampColumn + " DESC",
	).Limit(
		1,
	)

	err = sqrl.QueryRowWith(stdlibQueryRow(tx.QueryRow), q).Scan(
		&stressed, &flare,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var now pgtype.Timestamptz
	now.Time = time.Now()
	now.Status = pgtype.Present
	_, err = s.sb.Insert(
		internal.EventsTable,
	).Values(
		&now, req.GetPainLevel().String(), stressed, flare,
	).RunWith(
		tx,
	).Exec()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *Server) ToggleFlare(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var stressed bool
	var flare bool
	var painLevel string
	q := s.sb.Select(
		internal.PainLevelColumn,
		internal.StressedColumn,
		internal.FlareColumn,
	).From(
		internal.EventsTable,
	).OrderBy(
		internal.TimestampColumn + " DESC",
	).Limit(
		1,
	)

	err = sqrl.QueryRowWith(stdlibQueryRow(tx.QueryRow), q).Scan(
		&painLevel, &stressed, &flare,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var now pgtype.Timestamptz
	now.Time = time.Now()
	now.Status = pgtype.Present
	_, err = s.sb.Insert(
		internal.EventsTable,
	).Values(
		&now, painLevel, stressed, !flare,
	).RunWith(
		tx,
	).Exec()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *Server) ToggleStress(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var stressed bool
	var flare bool
	var painLevel string
	q := s.sb.Select(
		internal.PainLevelColumn,
		internal.StressedColumn,
		internal.FlareColumn,
	).From(
		internal.EventsTable,
	).OrderBy(
		internal.TimestampColumn + " DESC",
	).Limit(
		1,
	)

	err = sqrl.QueryRowWith(stdlibQueryRow(tx.QueryRow), q).Scan(
		&painLevel, &stressed, &flare,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var now pgtype.Timestamptz
	now.Time = time.Now()
	now.Status = pgtype.Present
	_, err = s.sb.Insert(
		internal.EventsTable,
	).Values(
		&now, painLevel, !stressed, flare,
	).RunWith(
		tx,
	).Exec()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *Server) ToggleNoPain(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var stressed bool
	q := s.sb.Select(
		internal.StressedColumn,
	).From(
		internal.EventsTable,
	).OrderBy(
		internal.TimestampColumn + " DESC",
	).Limit(
		1,
	)

	err = sqrl.QueryRowWith(stdlibQueryRow(tx.QueryRow), q).Scan(
		&stressed,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var now pgtype.Timestamptz
	now.Time = time.Now()
	now.Status = pgtype.Present
	_, err = s.sb.Insert(
		internal.EventsTable,
	).Values(
		&now, proto.PainLevel_NO_PAIN.String(), stressed, false,
	).RunWith(
		tx,
	).Exec()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}

func (s *Server) GetEvents(req *proto.GetEventsRequest, srv proto.Monitor_GetEventsServer) error {
	q := s.sb.Select(
		internal.TimestampColumn,
		internal.PainLevelColumn,
		internal.StressedColumn,
		internal.FlareColumn,
	).From(
		internal.EventsTable,
	)

	if req.GetStart().GetSeconds() > 0 || req.GetStart().GetNanos() > 0 {
		var startTime pgtype.Timestamptz
		startTime.Set(time.Unix(req.GetStart().GetSeconds(), int64(req.GetStart().GetNanos())))
		buf, _ := startTime.EncodeText(nil, nil)
		q = q.Where(sqrl.GtOrEq{
			internal.TimestampColumn: string(buf),
		})
	}

	if req.GetEnd().GetSeconds() > 0 || req.GetEnd().GetNanos() > 0 {
		var endTime pgtype.Timestamptz
		endTime.Set(time.Unix(req.GetEnd().GetSeconds(), int64(req.GetEnd().GetNanos())))
		buf, _ := endTime.EncodeText(nil, nil)
		q = q.Where(sqrl.LtOrEq{
			internal.TimestampColumn: string(buf),
		})
	}

	rows, err := q.RunWith(s.db).QueryContext(srv.Context())
	if err != nil {
		return err
	}

	for rows.Next() {
		var event proto.Event
		var when pgtype.Timestamptz
		var painLevel string
		err = rows.Scan(&when, &painLevel, &event.Stressed, &event.Flare)
		if err != nil {
			return err
		}

		event.Timestamp = &timestamp.Timestamp{
			Seconds: when.Time.Unix(),
			Nanos:   int32(when.Time.UnixNano()),
		}
		event.PainLevel = proto.PainLevel(proto.PainLevel_value[painLevel])

		err = srv.Send(&event)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
