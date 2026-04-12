import java.io.*;
import java.net.Socket;

public class TestClient {
    public static void main(String[] args) throws Exception {
        Socket socket = new Socket("127.0.0.1", 6767);
        OutputStream out = socket.getOutputStream();

        System.out.println("Connected! Sending tap space...");

        // Send: action=2 (tap), keyID=0x20 (space), duration=0
        byte[] packet = new byte[4];
        packet[0] = 2;      // Tap
        packet[1] = 0x01;   // Left Click
        packet[2] = 0;      // Duration high byte
        packet[3] = 0;      // Duration low byte

        out.write(packet);
        out.flush();

        System.out.println("Packet sent!");

        Thread.sleep(500);
        socket.close();
    }
}

