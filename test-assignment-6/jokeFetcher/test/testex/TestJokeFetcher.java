package testex;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.runners.MockitoJUnitRunner;

import java.util.Arrays;
import java.util.Date;
import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.mockito.Mockito.*;
import static org.hamcrest.CoreMatchers.*;

import static junit.framework.TestCase.assertEquals;

@RunWith(MockitoJUnitRunner.class)
public class TestJokeFetcher {

    @Test
    public void testGetAvailableTypes() {
        IDateFormatter df = mock(IDateFormatter.class);
        JokeFetcher jf = new JokeFetcher(df);
        List<String> types = jf.getAvailableTypes();
        assertThat(types, hasItems("eduprog","chucknorris","moma","tambal"));
        assertEquals(5, types.size());
    }

    @Test
    public void testIsStringValid() {
        IDateFormatter df = mock(IDateFormatter.class);
        JokeFetcher jf = new JokeFetcher(df);
        assertEquals(false, jf.isStringValid("some,string"));
        assertEquals(true, jf.isStringValid("chucknorris"));
    }

    @Test
    public void TestGetJokes() throws JokeException {
        IDateFormatter df = mock(IDateFormatter.class);
        /*
        the problem with this is that the function can call four different private methods depending on the function
        parameters. It also both fetches the jokes, creates joke and jokes objects etc.
         */
        JokeFetcher jf = new JokeFetcher(df);
        Jokes jokes = jf.getJokes("chucknorris", "Europe/Copenhagen");
        assertThat(jokes, notNullValue());
        verify(df,  times(1)).getFormattedDate("Europe/Copenhagen", new Date());
        System.out.println(jokes.getJokes());
    }
}
