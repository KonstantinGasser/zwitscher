package org.bdea.zwitscher;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;

public class ProfanityFilter {

    private static Set<String> words = new HashSet<>();
    private static int largestWordLength = 0;
    private static boolean initialized = false;

    private static void initialize() {
        try {
            BufferedReader reader = new BufferedReader(new FileReader("swearwords.txt"));
            String word = "";
            while ((word = reader.readLine()) != null) {
                word = word.replaceAll(" ", "");
                int length = word.length();
                if(length == 0) continue;
                if(length > largestWordLength) {
                    largestWordLength = word.length();
                }
                words.add(word);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        initialized = true;
    }

    public static boolean containsProfanity(String input) {
        if (input == null) return false;
        if (!initialized) initialize();

        // prepare input
        input = input.replaceAll("1","i");
        input = input.replaceAll("!","i");
        input = input.replaceAll("3","e");
        input = input.replaceAll("4","a");
        input = input.replaceAll("@","a");
        input = input.replaceAll("5","s");
        input = input.replaceAll("7","t");
        input = input.replaceAll("0","o");
        input = input.replaceAll("9","g");
        input = input.toLowerCase().replaceAll("[^a-zA-Z]", "");

        /*
        All of this complexity is okay by the way, because it only has to scale with the length of a single tweet,
        not with the amount of tweets.
         */
        for (int start = 0; start < input.length(); start++) {
            for (int offset = 1; offset < (input.length()+1 - start) && offset < largestWordLength; offset++)  {
                String wordToCheck = input.substring(start, start + offset);
                if (words.contains(wordToCheck)) {
                    return true;
                }
            }
        }
        return false;
    }
}
