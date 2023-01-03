<template>
    <div class="login_view">
        <div class="login_box">
            <div>
                <h2 class="login_show">登录</h2>
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
                    <el-button type="info" @click="regist">注册</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script>
import axios from "axios"
import config from '../../config/config'

export default {
    data() {
        return {
            loginData: {
                username: "",
                password: ""
            }
        }
    },
    methods: {
        login() {
            axios.post(config.host + "/api/user/login", this.loginData).then(res => {
                if (res.data.code === 0) {
                    localStorage.setItem("user", res.data.data.username)
                    axios.defaults.headers.common['Authorization'] = res.data.data.token
                    this.$router.push("/")
                }
            })
        },
        regist(){

        }
    }
}
</script>

<style lang="less" scoped>
.login_view{
    background-color: #2b4b6b;
    height: 100%;
}

.login_box{
    width: 450px;
    height: 300px;
    background-color: aliceblue;
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
}

.login_button{
    display: flex;
    justify-content: flex-end;
}
.login_form{
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
}

.login_show{
    position: absolute;
    width: 20%;
    left: 45%;
}
</style>