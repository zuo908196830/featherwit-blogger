import axios from "axios"

export default {
    setLoginStatus(user, token) {
        localStorage.setItem("user", user)
        axios.defaults.headers.common['Authorization'] = token
        localStorage.setItem("loginStatus", true)
        localStorage.setItem("token", token)
    }
}