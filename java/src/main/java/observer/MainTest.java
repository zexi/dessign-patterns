package observer;

public class MainTest {
    public static void main(String[] args) {
        WeatherData wd = new WeatherData();
        wd.addObserver(new CurrentConditionsDisplay(wd));
        wd.addObserver(new CurrentConditionsDisplay(wd));
        wd.setMeasurements(32, 14, 100);
    }
}
