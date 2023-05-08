import axios from "axios"

export default {
    setLoginStatus(user, token, headshot) {
        localStorage.setItem("user", user)
        axios.defaults.headers.common['Authorization'] = token
        localStorage.setItem("loginStatus", true)
        localStorage.setItem("token", token)
        localStorage.setItem("headshot", headshot)
    }
}