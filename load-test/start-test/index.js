import http from 'k6/http';
import {sleep} from 'k6'

export const options = {
    summaryTimeUnit: 's',
    duration: '5s',
    vus: 1000
    // stages: [
    //     {
    //         duration: "1s",
    //         target: 100
    //     },
    //     {
    //         duration: "1s",
    //         target: 500
    //     },
    //     // {
    //     //     duration: "1s",
    //     //     target: 1500
    //     // },
    //     // {
    //     //     duration: "1s",
    //     //     target: 2000
    //     // }
    // ]
}

export default function () {
    const payload = JSON.stringify({
        test_id: 1
    })
    const params = {
        headers: {
            'Authorization': `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2NzA4MTIzMjMsIm5iZiI6MTY3MDcyNTkyMywiaWF0IjoxNjcwNzI1OTIzLCJqdGkiOiIxIn0.gFmzUyH7OzAtrZgpxJNAtbZyu4qJM1GpgPPHHwl7FZ8`
        }
    }
    const response = http.post("http://localhost:21000/api/v1/answersheets/start", payload, params)

    // const response = http.get("http://localhost:21000/health")
    // console.log(response.body)

}