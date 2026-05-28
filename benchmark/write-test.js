import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 100,
  duration: '20s',
};

export default function () {

  const payload = JSON.stringify({
    name: `User ${Math.random()}`,
    email: `test${Math.random()}@gmail.com`,
    age: 25
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let res = http.post(
    'http://localhost:8083/users',
    payload,
    params
  );

  check(res, {
    'status is 201': (r) => r.status === 201,
  });

  sleep(1);
}