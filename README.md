# zwitscher
This a 3 day project for our lecture in Big-Data Engineering.
Idea of the project was it to come up with a solution which could handle heavy data read/write.
Further the solution shoudl support the following: 
- Some infrastructure supporting heavy read and write of data
- include streaming framework to do something fun with the data
- create personalized dashboards for user with tweets from their followers
- fetch tweets for user X
- post new tweet

# Solution
We use Kafka to transport data around in the infrastructre. There are topics pipeing raw tweets from the user to a Flink-Job which would filter out
tweets bad language. This job then pushs the "cleaned" tweet into a new topic where it will get picked up by yet another Flink-Job. Now as set by the task we should build a dashboard for different user. Our approach is it to have a streaming job which acts as a fanout. The Fanout would pickup each newly tweeted tweet and match it to all the tweet_user's followers. The job then updates the personalized dashboard record in the memcached-container. 
