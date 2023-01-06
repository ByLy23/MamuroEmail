import type { EmailInterface } from "@/interface/contentInterface";
import axios from "axios";
const apiStorage = axios.create({
  baseURL: "http://127.0.0.1:8080",
  headers: {
    "Content-Type": "application/json",
  },
});
export default {
  // getLocalStorage() {
  //   const data = localStorage.getItem("searchData");
  //   if (data) {
  //     const parsedData = JSON.parse(data);
  //     this.query = parsedData.phrase;
  //     this.contentSize = parsedData.page;
  //     this.totalPages = Math.ceil(parsedData.value / 20);
  //   }
  // },
  // addPage() {
  //   this.getLocalStorage();
  //   this.contentSize += 1;
  //   this.search(this.query, this.contentSize);
  // },
  // substractPage() {
  //   this.getLocalStorage();
  //   this.contentSize -= 1;
  //   this.search(this.query, this.contentSize);
  // },
  async search(phrase: string, pageNumber: number = 0) {
    const body = JSON.stringify({
      search_type: "matchphrase",
      query: {
        term: phrase,
        start_time: "2021-06-02T14:28:31.894Z",
        end_time: new Date().toISOString(),
      },
      from: pageNumber * 20 - 20,
      max_results: 20,
      _source: [],
    });
    const response = await apiStorage.post("/api/search", body);
    console.log(response.data.hits, response.status);
    if (response.status === 200) {
      return response.data;
    } else {
      return false;
    }
    // const response = await fetch("http://127.0.0.1:8080/api/search", options);
    // const data = await response.json();
    // console.log(data);
    // this.pruebaResponse = data.content;
  },
};
