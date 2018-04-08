package testex;

import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;
import testex.jokefetching.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Date;
import java.util.List;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.mockito.Mockito.*;
import static org.hamcrest.CoreMatchers.*;

import static junit.framework.TestCase.assertEquals;

@RunWith(MockitoJUnitRunner.class)
public class TestJokeFetcher {

    @Mock IDateFormatter df;
    @Mock IFetcherFactory ff;
    @Mock Moma moma;
    @Mock ChuckNorris chuckNorris;
    @Mock Edu edu;
    @Mock Tambal tambal;

    private JokeFetcher jf;

    @Before
    public void setup() {
        List<IJokeFetcher> fetchers = Arrays.asList(edu,chuckNorris,moma,tambal);
        when(ff.getJokeFetchers("EduJoke,ChuckNorris,Moma,Tambal")).thenReturn(fetchers);
        when(ff.getJokeFetchers("Moma")).thenReturn(Arrays.asList(moma));
        when(ff.getJokeFetchers("ChuckNorris")).thenReturn(Arrays.asList(chuckNorris));
        when(ff.getJokeFetchers("EduJoke")).thenReturn(Arrays.asList(edu));
        when(ff.getJokeFetchers("Tambal")).thenReturn(Arrays.asList(tambal));
        List<String> types = Arrays.asList("EduJoke","ChuckNorris","Moma","Tambal");
        when(ff.getAvailableTypes()).thenReturn(types);
        jf = new JokeFetcher(df, ff);
    }

    @Test
    public void testGetAvailableTypes() {
        List<String> types = ff.getAvailableTypes();
        assertThat(types, hasItems("EduJoke","ChuckNorris","Moma","Tambal"));
        assertEquals(4, types.size());
    }

    @Test
    public void testIsStringValid() {
        assertEquals(false, jf.isStringValid("some,string"));
        assertEquals(true, jf.isStringValid("ChuckNorris"));
    }

    @Test
    public void TestGetJokes() throws JokeException {
        /*
        the problem with this is that the function can call four different private methods depending on the function
        parameters. It also both fetches the jokes, creates joke and jokes objects etc.
         */
        Jokes jokes = jf.getJokes("ChuckNorris", "Europe/Copenhagen");
        assertThat(jokes, notNullValue());
        verify(df,  times(1)).getFormattedDate("Europe/Copenhagen", new Date());
        System.out.println(jokes.getJokes());
    }

    @Test
    public void TestGetAll() throws JokeException {
        Joke testJoke = new Joke("xd joke", "xd.xd");
        when(moma.getJoke()).thenReturn(testJoke);
        Jokes jokes = jf.getJokes("EduJoke,ChuckNorris,Moma,Tambal", "Europe/Copenhagen");
        assertThat(jokes.getJokes(), hasItems(testJoke));
        verify(ff, times(1)).getJokeFetchers("EduJoke,ChuckNorris,Moma,Tambal");
    }

    @Test
    public void TestGetMomaJoke() throws JokeException {
        Joke testJoke = new Joke("xd joke", "xd.xd");
        when(moma.getJoke()).thenReturn(testJoke);
        Jokes momaJokes = jf.getJokes("Moma", "Europe/Copenhagen");
        assertThat(momaJokes.getJokes(), hasItems(testJoke));
        verify(ff, times(1)).getJokeFetchers("Moma");
    }

    @Test
    public void TestGetChuckNorrisJoke() throws JokeException {
        Joke testJoke = new Joke("xd joke", "xd.xd");
        when(chuckNorris.getJoke()).thenReturn(testJoke);
        Jokes chuckNorrisJokes = jf.getJokes("ChuckNorris", "Europe/Copenhagen");
        assertThat(chuckNorrisJokes.getJokes(), hasItems(testJoke));
        verify(ff, times(1)).getJokeFetchers("ChuckNorris");
    }

    @Test
    public void TestGetEduJoke() throws JokeException {
        Joke testJoke = new Joke("xd joke", "xd.xd");
        when(edu.getJoke()).thenReturn(testJoke);
        Jokes eduJokes = jf.getJokes("EduJoke", "Europe/Copenhagen");
        assertThat(eduJokes.getJokes(), hasItems(testJoke));
        verify(ff, times(1)).getJokeFetchers("EduJoke");
    }

    @Test
    public void TestGetTambalJoke() throws JokeException {
        Joke testJoke = new Joke("xd joke", "xd.xd");
        when(tambal.getJoke()).thenReturn(testJoke);
        Jokes tambalJokes = jf.getJokes("Tambal", "Europe/Copenhagen");
        System.out.println(tambalJokes.getJokes());
        System.out.println(testJoke);
        assertThat(tambalJokes.getJokes(), hasItems(testJoke));
        verify(ff, times(1)).getJokeFetchers("Tambal");
    }
}
