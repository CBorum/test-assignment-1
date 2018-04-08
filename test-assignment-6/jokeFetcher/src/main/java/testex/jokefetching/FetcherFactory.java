package testex.jokefetching;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class FetcherFactory implements IFetcherFactory {

    private final List<String> availableTypes = Arrays.asList("EduJoke", "ChuckNorris", "Moma", "Tambal");

    @Override
    public List<String> getAvailableTypes(){
        return availableTypes;
    }

    @Override
    public List<IJokeFetcher> getJokeFetchers(String jokesToFetch) {
        List<IJokeFetcher> res = new ArrayList<>();
        String[] toFetch = jokesToFetch.split(",");
        for (String type : toFetch) {
            switch (type) {
                case "EduJoke":
                    res.add(new Edu());
                    break;
                case "ChuckNorris":
                    res.add(new ChuckNorris());
                    break;
                case "Moma":
                    res.add(new Moma());
                    break;
                case "Tambal":
                    res.add(new Tambal());
                    break;
            }
        }
        return res;
    }
}

