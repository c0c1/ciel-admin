import useSWR from "swr";
import {isEmpty} from "../utils/util";
import {KeyUserInfo} from "../consts/const";

const userFetcher = async () => {
    let u = localStorage.getItem(KeyUserInfo)
    if (isEmpty(u)) {
        let error = new Error("Not authorized!");
        error.status = 403
        throw error
    }
    return JSON.parse(u)
}
export const useUser = () => {
    const {data, mutate, error} = useSWR(`api_user`, userFetcher)
    const loading = !data && !error;
    const loggedOut = error && error.status === 403;
    return {loading, loggedOut, u: data, mutate}
}