package org.bdea.zwitscher;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class RandomId {
    private static List<Integer> ids;
    private static boolean initialized = false;

    private static void initialize() {
        ids = new ArrayList<>();
        try {
            BufferedReader reader = new BufferedReader(new FileReader("ids.txt"));
            String line;
            while((line = reader.readLine()) != null) {
                ids.add(Integer.parseInt(line));
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        initialized = true;
    }

    public static int get() {
        if (!initialized) initialize();
        return ids.get((int)Math.floor(Math.random() * 100));
    }
}
