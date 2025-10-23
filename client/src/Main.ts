import { createApp } from "vue";
import App from "./App.vue";
import router from "./Router.ts";
import { pinia } from "./Pinia.ts";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
    faCheck,
    faCheckCircle,
    faInfoCircle,
    faExclamationTriangle,
    faExclamationCircle,
    faArrowUp,
    faAngleRight,
    faAngleLeft,
    faAngleDown,
    faEye,
    faEyeSlash,
    faCaretDown,
    faCaretUp,
    faUpload,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(
    faCheck,
    faCheckCircle,
    faInfoCircle,
    faExclamationTriangle,
    faExclamationCircle,
    faArrowUp,
    faAngleRight,
    faAngleLeft,
    faAngleDown,
    faEye,
    faEyeSlash,
    faCaretDown,
    faCaretUp,
    faUpload,
);

const app = createApp(App);
app.component("VueFontAwesome", FontAwesomeIcon);
app.use(router);
app.use(pinia);
import Buefy from "buefy";
app.use(Buefy, { defaultIconComponent: "vue-font-awesome", defaultIconPack: "fas" });
app.mount("#app");
