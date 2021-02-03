package org.bdea.zwitscher;

import java.io.*;
import java.util.ArrayList;
import java.util.List;

public class RandomId {
    private static List<Integer> ids;
    private static boolean initialized = false;

    private static void initialize() throws IOException {
        ids = new ArrayList<>();
        // try {
//        ClassLoader cl = RandomId.class.getClassLoader();
        ClassLoader cl = Thread.currentThread().getContextClassLoader();
//        BufferedReader reader = new BufferedReader(new FileReader("src/main/resources/ids.txt"));

        InputStream is = cl.getResourceAsStream("ids.txt");
        BufferedReader reader = new BufferedReader(new InputStreamReader(is));
        String line;
        while((line = reader.readLine()) != null) {
            ids.add(Integer.parseInt(line));
        }
        // } catch (IOException e) {
        //     e.printStackTrace();
        // }
        initialized = true;
    }

    public static int get() throws IOException {
        if (!initialized) initialize();
        return ids.get((int)Math.floor(Math.random() * ids.size()));
    }
}
