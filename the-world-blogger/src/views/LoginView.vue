<template>
    <div class="login_view">
        <gvb-nav/>
        <gvb-banner/>
        <div class="login_box">
            <div>
                <h2 class="login_show">登录</h2>
            </div>
            <div style="position:absolute; top:20%; height:30px; width:100%;">
                <el-alert v-if="loginCode === 1001" title="该用户不存在" type="error" center show-icon closable>
                </el-alert>
                <el-alert v-if="loginCode === 1004" title="密码错误" type="error" center show-icon closable>
                </el-alert>
                <el-alert v-if="loginCode === 1003" title="格式错误" type="error" center show-icon closable>
                </el-alert>
                <el-alert v-if="loginCode === 500" title="服务端错误" type="error" center show-icon closable>
                </el-alert>
            </div>
            <el-form class="login_form" :model="loginData">
                <el-form-item>
                    <el-input prefix-icon="el-icon-user" v-model="loginData.username"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-input prefix-icon="el-icon-lock" v-model="loginData.password" type="password"></el-input>
                </el-form-item>
                <el-form-item class="login_button">
                    <el-button type="primary" @click="login">登录</el-button>
                    <router-link to="/register">
                        <el-button type="info" @click="regist">注册</el-button>
                    </router-link>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
import axios from "axios"
import LoginFunc from "../js/LoginFunc"
import gvbNav from "@/views/components/gvbNav";
import gvbBanner from "@/views/components/gvbBanner";

export default {
    components: {
        gvbNav,
        gvbBanner,
    },
    data() {
        return {
            loginData: {
                username: "",
                password: ""
            },
            loginCode: 0
        }
    },
    methods: {
        login() {
            axios.post("/api/user/login", this.loginData).then(res => {
                console.log(res.status)
                if (res.status === 200) {
                    this.loginCode = res.data.code
                    if (this.loginCode === 0) {
                        if (res.data.data.nickname !== "") {
                            LoginFunc.setLoginStatus(res.data.data.nickname, res.data.data.token, res.data.data.headshot)
                        } else {
                            LoginFunc.setLoginStatus(res.data.data.username, res.data.data.token, res.data.data.headshot)
                        }
                        this.$store.state.username = res.data.data.username
                        this.$router.push("/")
                    }
                } else {
                    // todo: 登陆失败的对应操作
                    this.loginCode = 500
                }
            })
        },
        regist() {

        },
        getStatus() {
            alert(localStorage.getItem('loginStatus'))
        }
    }
}
</script>

<style lang="less" scoped>
.login_view {
    background-color: #2b4b6b;
    height: 100%;
}

.login_box {
    width: 450px;
    height: 300px;
    background-color: aliceblue;
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
}

.login_button {
    display: flex;
    justify-content: flex-end;
}

.login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
}

.login_show {
    position: absolute;
    width: 20%;
    left: 45%;
}
</style>