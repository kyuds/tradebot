import time
from confluent_kafka import Producer, Consumer

def main():
    print("Starting predictor.py...")
    # sleep 15 seconds for AdminClient and Kafka delays
    time.sleep(15)

    producer, consumer = None, None

    try:
        producer = Producer({
            'bootstrap.servers': 'kafka-1:29092'
        })
        consumer = Consumer({
            'bootstrap.servers': 'kafka-1:29092',
            'group.id': 'group',
            'auto.offset.reset': 'latest'
        })
    except Exception as e:
        print("Problem with creating producers or consumers")
        print(e)
    
    if not producer or not consumer:
        return

    print("Created consumer and producer")


if __name__ == "__main__":
    main()
    while True:
        time.sleep(5)
        print("hihi")