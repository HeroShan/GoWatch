package ws

import (
    "log"
    "github.com/streadway/amqp"
)



func MqConn()(chh *amqp.Channel){
    conn, err := amqp.Dial("amqp://admin:shanhuijie@47.104.225.152:5672")
    ch, err := conn.Channel();if err != nil{
        log.Print(err)
    }
    chh = ch
    return chh
}

func Send(ch *amqp.Channel,queuename string,body []byte) bool {
    defer ch.Close()
    q, err := ch.QueueDeclare(
        queuename, //Queue name
        true, //durable
        false,
        false,
        false,
        nil,
    )
    err = ch.Publish(
        "",     //exchange
        q.Name, //routing key(queue name)
        false,
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent, //Msg set as persistent
            ContentType:  "text/plain",
            Body:         []byte(body),
        })
    if err !=nil{
        log.Print(err)
        return false
    }
    return true
}

func Receive(ch *amqp.Channel,queuename string) {
    q, err := ch.QueueDeclare(
        queuename, //Queue name
        true, //durable
        false,
        false,
        false,
        nil,
    )
    if err != nil{
        log.Printf("connect a message: %v\n", err)
    }
    msgs, _ := ch.Consume(
        q.Name,
        "",
        true,  //Auto Ack
        false,
        false,
        false,
        nil,
    )
	go func(){
        for msg := range msgs {
            log.Println(string(msg.Body))
        }
    }()
}

