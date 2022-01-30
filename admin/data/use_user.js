import {userInfo} from "../api/sys";
import useSWR from "swr";

export default function use_user() {
    const {data, mutate, error} = useSWR("api_user", userInfo)
    const loading = !data && !error;
    const loggedOut = error && error.status === 403;
    return {loading, loggedOut, user: data, mutate}
}