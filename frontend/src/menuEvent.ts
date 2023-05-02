import {EventsOn} from "../wailsjs/runtime"
// import { useRouter } from "vue-router";
import { router } from "./router";

// const route = useRouter();

// List Connection Events
EventsOn("LIST_CONNECTIONS",() => {
    router.push("/connections");
})

EventsOn("REFRESH_CONNECTION", () => {
    router.push("/home")
})