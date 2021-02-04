package org.bdea.zwitscher;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

public class RandomId {
    private static List<String> ids;
    private static boolean initialized = false;

    private static void initialize() throws IOException {
        ids = new ArrayList<>();
        ClassLoader cl = Thread.currentThread().getContextClassLoader();

        InputStream is = cl.getResourceAsStream("ids.txt");
        BufferedReader reader = new BufferedReader(new InputStreamReader(is));
        String line;
        while((line = reader.readLine()) != null) {
            ids.add(line);
        }
        initialized = true;
    }

    public static String get() throws IOException {
        if (!initialized) initialize();
        return ids.get((int)Math.floor(Math.random() * ids.size()));
    }
}
