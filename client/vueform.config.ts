import en from "@vueform/vueform/locales/en";
import vueform from "@vueform/vueform/dist/vueform";
import { defineConfig } from "@vueform/vueform";
import axios from "axios";
import "@vueform/vueform/dist/vueform.css";

axios.defaults.headers.put = {
  "Content-Type": "application/json",
};
axios.defaults.headers.post = {
  "Content-Type": "application/json",
};

export default defineConfig({
  theme: vueform,
  locales: { en },
  locale: "en",
});
