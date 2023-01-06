import axios from "axios";
const apiStorage = axios.create({
  baseURL: "http://127.0.0.1:8080",
  headers: {
    "Content-Type": "application/json",
  },
});
export const getSearch = async (phrase: string, pageNumber = 1) => {
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
  // console.log(response.data.hits, response.status);
  if (response.status === 200) {
    return response.data;
  } else {
    return false;
  }
};
