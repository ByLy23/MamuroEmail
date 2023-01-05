import type {
  contentInterface,
  EmailInterface,
} from "@/interface/contentInterface";
import { defineComponent } from "vue";
import axios from "axios";
const apiStorage = axios.create({
  baseURL: "http://127.0.0.1:8080",
  headers: {
    "Content-Type": "application/json",
  },
});
export default defineComponent({
  name: "SearchComponent",
  setup() {
    return {
      pruebaResponse: [] as EmailInterface[],
      fetching: false,
    };
  },
  methods: {
    async search(phrase: string, pageNumber: number) {
      this.fetching = true;
      const options = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          search_type: "matchphrase",
          query: {
            term: phrase,
            start_time: "2021-06-02T14:28:31.894Z",
            end_time: new Date().toISOString(),
          },
          from: pageNumber,
          max_results: 20,
          _source: [],
        }),
      };
      const body = JSON.stringify({
        search_type: "matchphrase",
        query: {
          term: phrase,
          start_time: "2021-06-02T14:28:31.894Z",
          end_time: new Date().toISOString(),
        },
        from: pageNumber,
        max_results: 20,
        _source: [],
      });
      const response = await apiStorage.post("/api/search", body);
      // const response = await fetch("http://127.0.0.1:8080/api/search", options);
      // const data = await response.json();
      // console.log(data);
      // this.pruebaResponse = data.content;
      this.fetching = false;
    },
  },
  async mounted() {
    await this.search("test", 0);
  },
});
