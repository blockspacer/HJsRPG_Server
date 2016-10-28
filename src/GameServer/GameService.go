package GameServer

import (
    "os"
    "time"
    "fmt"
    "net"
    "strconv"
)

// ---------------------------------------------------------

// GameService : 게임 서버의 전체적인 처리를 하는 게임 서비스이다.
type GameService struct {
    logOutput *os.File

    port int

    connectClientCh chan net.Conn
}

// Start : 게임 서비스를 실행한다.
func (g *GameService) Start() {
    
    // Log를 남길 파일 생성
    now := time.Now().Format("2006-01-02 15h 04m 05s")
    
    os.Mkdir("Log", 0777)
    o, err := os.Create("Log/Log(" + now + ").log")
    if err != nil {
        panic(err)
    }
    g.logOutput = o

    // 클라이언트 Listen 열기
    go g.ListenClient()
    // 메인 루프 실행
    go g.Loop()

    g.WriteLog("GameService Start.")
}

// WriteLog : 로그파일에 로그를 남긴다.
func (g *GameService) WriteLog(log string) {
    now := time.Now().Format("2006-01-02 15h 04m 05s")

    fmt.Println(now + ": " + log)
    _, err := g.logOutput.WriteString(now + ": " + log + "\n");
    if err != nil {
        panic(err)
    }
}

// ListenClient : 클라이언트 접속을 기다린다.
func (g *GameService) ListenClient() {
    // 설정한 Port로 Listen 한다.
    listen, err := net.Listen("tcp4", ":" + strconv.Itoa(g.port))
    defer listen.Close()
    if err != nil {
        g.WriteLog(fmt.Sprintf("Fail to listen on port %d.", g.port))
       // g.WriteLog(err.(string))
        return
    }
    g.WriteLog("Start to listen to connect Client...")

    for {
        conn, err := listen.Accept()
        if err != nil {
            g.WriteLog("Fail to accept the client.")
        } else {
            g.connectClientCh <- conn
        }
    }
}

// Loop : GameService의 메인 루프문
func (g *GameService) Loop() {/*
LOOP:
   for {
        select {
            case conn := <-g.connectClientCh: // 클라이언트 연결 됨
                g.WriteLog("Client connected.")
                conn.
        }
    }*/
}

// ---------------------------------------------------------

// ConnectionInfo : 연결된 클라이언트의 정보이다
type ConnectionInfo struct {
    readCh chan uint
    sendCh chan uint
    
    conn net.Conn
}

// ReadLoop : 들어오는 정보를 받아내는 루프문
func (c *ConnectionInfo) ReadLoop() {

}

// SendLoop : 클라이언트로 정보를 보내는 루프문
func (c *ConnectionInfo) SendLoop() {

}

// ---------------------------------------------------------


// CreateGameService : 게임 서비스를 생성한다.
func CreateGameService(port int) *GameService {
    g := GameService {
        port: port,
        connectClientCh: make(chan net.Conn),
    }

    return &g
}

// NewConnectionInfo : 클라이언트 연결에 관한 새로운 정보 생성
func NewConnectionInfo(conn net.Conn) *ConnectionInfo {
    c := ConnectionInfo {
        conn: conn,
        readCh: make(chan uint, 1),
        sendCh: make(chan uint, 1),
    }

    go c.ReadLoop();
    go c.SendLoop();

    return &c;
}