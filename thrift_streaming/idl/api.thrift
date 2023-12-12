namespace go echo

struct EchoRequest {
    1: required string message,
}

struct EchoResponse {
    1: required string message,
}

exception EchoException {
  1: string message
}


service EchoService {
    // streaming api (HTTP2)
    EchoResponse EchoBidirectional (1: EchoRequest req1) (streaming.mode="bidirectional"),
    EchoResponse EchoClient (1: EchoRequest req1) (streaming.mode="client"),
    EchoResponse EchoServer (1: EchoRequest req1) (streaming.mode="server"),
    EchoResponse EchoUnary (1: EchoRequest req1) (streaming.mode="unary"),

    // KitexThrift
    EchoResponse EchoPingPong (1: EchoRequest req1, 2: EchoRequest req2) throws (1: EchoException e),
    void EchoOneway(1: EchoRequest req1),
    void Ping(),
}
