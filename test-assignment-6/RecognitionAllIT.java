package net.sf.javaanpr.test;

import junit.framework.TestSuite;
import net.sf.javaanpr.imageanalysis.CarSnapshot;
import net.sf.javaanpr.intelligence.Intelligence;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;
import org.junit.runners.Parameterized.Parameter;
import org.junit.runners.Parameterized.Parameters;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.xml.sax.SAXException;

import javax.xml.parsers.ParserConfigurationException;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.Properties;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.core.IsEqual.equalTo;
import static org.hamcrest.core.IsNull.notNullValue;
import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertTrue;

@RunWith(Parameterized.class)
public class RecognitionAllIT {
    private static final Logger logger = LoggerFactory.getLogger(RecognitionAllIT.class);

    @Parameters(name = "{index}: - plate: {1}")
    public static Collection<Object[]> data() throws IOException {
        String snapshotDirPath = "src/test/resources/snapshots";
        String resultsPath = "src/test/resources/results.properties";
        InputStream resultsStream = new FileInputStream(new File(resultsPath));

        Properties properties = new Properties();
        properties.load(resultsStream);
        resultsStream.close();
        assertTrue(properties.size() > 0);

        File snapshotDir = new File(snapshotDirPath);
        File[] snapshots = snapshotDir.listFiles();
        assertNotNull(snapshots);
        assertTrue(snapshots.length > 0);

        ArrayList<Object[]> objects = new ArrayList<>();

        for (File f : snapshots) {
            String plateCorrect = properties.getProperty(f.getName());
            objects.add(new Object[]{f, plateCorrect});
        }

        return objects;
    }

    @Parameter
    public File inputFile;

    @Parameter(1)
    public String plateCorrect;

    private Intelligence intel;

    @Before
    public void before() throws IOException, SAXException, ParserConfigurationException {
        intel = new Intelligence();
    }

    @Test
    public void intelligenceSingleTest() throws IOException {
        CarSnapshot carSnap = new CarSnapshot(new FileInputStream(inputFile));
        assertThat(carSnap, notNullValue());
        assertThat(carSnap.getImage(), notNullValue());
        assertThat(plateCorrect, notNullValue());

        String numberPlate = intel.recognize(carSnap, false);

        assertThat(numberPlate, notNullValue());
        logger.debug("snapname + result: " + plateCorrect + " - " + numberPlate);
        assertThat(plateCorrect, equalTo(numberPlate));
        carSnap.close();
    }
}
