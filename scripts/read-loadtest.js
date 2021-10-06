import http from "k6/http";
import { check } from "k6";
const keywords = getKeywords();

export const options = {
  vus: 12,
  duration: "2m",
  insecureSkipTLSVerify: true,
  ext: {
    loadimpact: {
      projectID: 3554719,
      name: "read-loadtest",
    },
  },
};

export default function () {
  const startTimestamp = Date.now();
  const keyword = keywords[Math.floor(Math.random() * keywords.length)];
  const res = http.get(`http://localhost:3000/read?keyword=${keyword}`);
  check(res, {
    "response code was 200": (res) => res.status == 200,
  });
  console.log(
    JSON.stringify({
      status: res.status,
      duration: Date.now() - startTimestamp,
    })
  );
}

function getKeywords() {
  return [
    "perspiciatis",
    "iusto",
    "dignissimos",
    "fugiat",
    "repudiandae",
    "voluptatem",
    "autem",
    "corporis",
    "accusantium",
    "totam",
  ];
}
