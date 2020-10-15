package me.dennis;

public class Counter {

    private int count;

    public Counter() {
        this.count = 0;
    }

    public int getCount() {
        return this.count;
    }

    public void countUp() {
        this.count++;
    }

    public void countDown() {
        this.count--;
    }
}
