import http from "k6/http";
import { check } from "k6";

export default function () {
  const startTimestamp = Date.now();
  const res = http.get("http://localhost:3000/read?keyword=perspiciatis");
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
