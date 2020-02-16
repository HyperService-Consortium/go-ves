package sessionservice

import (
	logger2 "github.com/Myriad-Dreamin/go-ves/lib/log"
	"github.com/Myriad-Dreamin/go-ves/lib/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"log"
	"testing"
)

var (
	sessionIDNotFound                            = []byte("xx")
	sessionIDFindTransactionError                = []byte("xy")
	sessionIDPushTransactionNotNil               = []byte("xz")
	sessionIDAttestationSendErrorNotOk           = []byte("yx")
	sessionIDFindError                           = []byte("yy")
	sessionIDAttestationSendErrorNotOk2          = []byte("yz")
	sessionIDAttestationSendError                = []byte("zz")
	sessionIDDeserializeTransactionError         = []byte("xyz")
	sessionIDFindSessionWithAcknowledgeError     = []byte("x")
	sessionIDFindSessionWithGetAcknowledgedError = []byte("y")
)

type fields struct {
	cfg            *config.ServerConfig
	key            string
	accountDB      control.SessionAccountDBI
	db             control.SessionDBI
	sesFSet        control.SessionFSetI
	opInitializer  control.OpIntentInitializerI
	signer         control.Signer
	logger         types.Logger
	cVes           control.CentralVESClient
	respAccount    control.Account
	storage        control.SessionKV
	storageHandler control.StorageHandler
	dns            control.ChainDNS
	nsbClient      control.NSBClient
}

func MockSessionDB(ctl *gomock.Controller) *mock.SessionDB {
	return mock.NewSessionDB(ctl)
}

func MockSessionFSet(ctl *gomock.Controller) *mock.SessionFSet {
	return mock.NewSessionFSet(ctl)
}

func MockSessionAccountDB(ctl *gomock.Controller) *mock.SessionAccountDB {
	return mock.NewSessionAccountDB(ctl)
}

func MockCentralVESClient(ctl *gomock.Controller) *mock.CentralVESClient {
	return mock.NewCentralVESClient(ctl)
}

func createField(options ...interface{}) fields {
	ensureTestLogger()
	f := fields{
		logger: testLogger,
		cfg:    config.Default(),
	}
	for i := range options {
		switch o := options[i].(type) {
		case *mock.SessionDB:
			f.db = o
		case *mock.SessionFSet:
			f.sesFSet = o
		case *mock.SessionAccountDB:
			f.accountDB = o
		case *mock.CentralVESClient:
			f.cVes = o
		}
	}

	return f
}

func ensureTestLogger() {
	if testLogger == nil {
		if testing.Verbose() {
			var err error
			testLogger, err = logger.NewZapLogger(
				logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
			if err != nil {
				log.Fatal("init vesLogger error", "error", err)
			}
		} else {
			testLogger = logger2.NewNopLogger()
		}
	}
}

var testLogger logger.Logger

func checkErrorCode(t *testing.T, err error, i int) {
	t.Helper()
	if i != types.CodeOK {
		if f, ok := wrapper.FromError(err); ok {
			if f.GetCode() != i {
				t.Errorf("not expected code, error code %v, wantCode %v", f.GetCode(), i)
			} else {
				ensureTestLogger()
				testLogger.Info("expected good error", "error", err)
			}
		} else {
			t.Error("not frame error wrapped")
		}
	}
}