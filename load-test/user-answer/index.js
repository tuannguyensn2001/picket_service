import http from 'k6/http'

export const options = {
    duration: "5s",
    vus: 10
}

export default function () {
    const payload = JSON.stringify({
        question_id: 1,
        test_id: 1,
        answer: "A",
        previous_answer: "B"
    })
    const params = {
        headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0dWFubmd1eWVuc24yMDAxYUBnbWFpbC5jb20iLCJleHAiOjE2NzMzNjI4MTcsIm5iZiI6MTY3MzI3NjQxNywiaWF0IjoxNjczMjc2NDE3LCJqdGkiOiIxIn0.LKfHzbidbGOeHmsuhtQJGOyq5IGN_VudRh-QP7WBj3Q'
        }
    }
    http.post("http://localhost:21000/api/v1/answersheets/answer",payload,params)
}