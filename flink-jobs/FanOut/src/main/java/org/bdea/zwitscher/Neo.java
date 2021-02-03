package org.bdea.zwitscher;

import org.neo4j.driver.*;

import java.util.List;

public class Neo {
    private static Driver driver = null;
    private static String query = "MATCH (x:User)-[:FOLLOWS]->(:User {id: \"";
    private static String query2 = "\"}) RETURN x";

    public static List<String> getFollowers(final int id) {
        if (driver == null)
            driver = GraphDatabase.driver("bolt://neo4j:7687", AuthTokens.basic("neo4j", "sink"));

        try (Session session = driver.session()) {
            return session.run(query + id + query2).list(record -> record.get(0).asString());
        }
    }
}