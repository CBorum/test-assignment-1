package testex;

import org.junit.Test;

import java.util.Calendar;
import java.util.Date;
import java.util.GregorianCalendar;
import java.util.TimeZone;

import static org.junit.Assert.assertEquals;

public class TestDateFormatter {

    IDateFormatter df = new DateFormatter();

    @Test(expected = JokeException.class)
    public void testGetFormattedDate() throws JokeException {
        df.getFormattedDate("abc", new Date());
    }

    @Test
    public void testGetFormattedDate2() throws JokeException {
        Calendar date = new GregorianCalendar(2018, 3, 19);
        String timeZone = "Europe/Kiev";
        String abc = df.getFormattedDate(timeZone, date.getTime());
        TimeZone tz = TimeZone.getTimeZone(abc);
        assertEquals("GMT", tz.getID());
        assertEquals("19 Apr 2018 01:00 AM", abc);
    }
}
