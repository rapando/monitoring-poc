import http from 'k6/http';
import {check, sleep} from 'k6';
import {Trend} from 'k6/metrics';

// defin the custom metrics
const homeTrend = new Trend('get_home_response_time');
const countTrend = new Trend('get_count_response_time');
const addTrend = new Trend('get_add_response_time');

const BASE_URL = 'https://monitoring-api.bg.co.ke';

export let options = {
    stages: [
        {duration: '5s', target: 50},
        {duration: '5s', target: 100},
        {duration: '5s', target: 200},
        {duration: '5s', target: 300},
        {duration: '5s', target: 500},
        {duration: '30s', target: 200},
        {duration: '10s', target: 10},
        {duration: '10s', target: 10},
        {duration: '10s', target: 300},
        {duration: '10s', target: 500},
        {duration: '10s', target: 600},
        {duration: '10s', target: 1000},
    ],
    thresholds: {
        'get_home_response_time': ['p(95)<1000'],
        'get_count_response_time': ['p(95)<1000'],
        'get_add_response_time': ['p(95)<1000'],
    },
};

export default function() {
    let homeRes = http.get(`${BASE_URL}/`);
    check(homeRes, {'status was 200': (r) => r.status === 200});
    homeTrend.add(homeRes.timings.duration);

    let addRes = http.post(`${BASE_URL}/data`);
    check(addRes, {'status was 201': (r) => r.status === 201});
    addTrend.add(addRes.timings.duration);

    let countRes = http.get(`${BASE_URL}/data`);
    check(countRes, {'status was 200': (r) => r.status === 200});
    countTrend.add(countRes.timings.duration);

    sleep(1)
}


// ZYKkiPLGvCtm2jweG19BWE0tqf9PxIBFs75URvF5Q0QjpUZaJywEtVhpMQohrAxlz8qPX0XXXutxjx-Lso7Gww==