package org.bdea.zwitscher;

import net.spy.memcached.AddrUtil;
import net.spy.memcached.ConnectionFactoryBuilder;
import net.spy.memcached.MemcachedClient;
import net.spy.memcached.auth.AuthDescriptor;
import net.spy.memcached.auth.PlainCallbackHandler;

import java.io.IOException;
import java.util.Collections;

public class Memcached {
    private static MemcachedClient client = null;

    private static void initialize() throws IOException {
        AuthDescriptor ad = new AuthDescriptor(new String[] {"PLAIN"},
                new PlainCallbackHandler("username", "password"));

//        try {
            MemcachedClient mc = new MemcachedClient(
                    new ConnectionFactoryBuilder()
                            .setProtocol(ConnectionFactoryBuilder.Protocol.BINARY)
                            .setAuthDescriptor(ad).build(),
                    AddrUtil.getAddresses(Collections.singletonList("memcached:11211")));
//        } catch (IOException e) {
//             handle exception
//        }
    }

    public static void updateTimeline(String userId, String tweet) throws IOException {
        if (client == null) initialize();
        client.append(userId, ","+tweet);
    }
}
