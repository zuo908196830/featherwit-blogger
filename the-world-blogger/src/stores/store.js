import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex)

export const store = new Vuex.Store({
    state: {
        tagId: 0,
        searchName: "",
        loginStatus: false,
        username: "",
        token: "",
        headshot:"",
    }
})