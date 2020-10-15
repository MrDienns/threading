package me.dennis;

public class Threading {

    private final Counter counter = new Counter();

    public static void main(String[] args) throws InterruptedException {

        Threading threading = new Threading();

        Thread up = new Thread(threading::countUp);
        Thread down = new Thread(threading::countDown);

        up.start();
        down.start();

        up.join();
        down.join();

        System.out.println("Count: " + threading.counter.getCount());
    }

    private void countUp() {
        for (int i = 0; i < 1_000_000; i++) {
            this.counter.countUp();
        }
    }

    private void countDown() {
        for (int i = 0; i < 1_000_000; i++) {
            this.counter.countDown();
        }
    }
}
