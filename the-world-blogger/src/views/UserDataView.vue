<template>
    <div>
        <div>
            <GvbNav></GvbNav>
        </div>
        <div style="height: 70px;width: 100%;"></div>
        <div>
            <el-descriptions class="margin-top" title="个人信息" :column="3" :size="size" border>
                <template slot="extra">
                    <el-button type="primary" size="small">修改</el-button>
                </template>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-user"></i>
                        账号
                    </template>
                    {{ userData.username }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-mobile-phone"></i>
                        手机号
                    </template>
                    {{ userData.telephone }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-s-custom"></i>
                        昵称
                    </template>
                    {{ userData.nickname }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-mail"></i>
                        邮箱
                    </template>
                    {{ userData.mail }}
                </el-descriptions-item>
                <el-descriptions-item>
                    <template slot="label">
                        <i class="el-icon-office-tickets"></i>
                        简介
                    </template>
                    {{ userData.profile }}
                </el-descriptions-item>
            </el-descriptions>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import GvbNav from './components/gvbNav.vue'
import gvbNav from './components/gvbNav.vue'

export default {
    comments: {
        gvbNav
    },
    data() {
        return {
            userData: {
                username: "",
                role: 0,
                nickname: "",
                telephone: "",
                mail: "",
                profile: "",
            }
        };
    },
    methods: {
        getUser() {
            axios.defaults.headers.common["Authorization"] = localStorage.getItem("token");
            axios.get("/api/user/data").then(res => {
                if (res.data.code === 0) {
                    this.userData = res.data.data;
                }
                else if (res.data.code === 1006) {
                    // todo 提示未登录
                    this.$message({
                        message: "您还未登录",
                        type: "warning"
                    });
                    this.$router.push("/login");
                }
            }).catch(() => {
            });
        }
    },
    created() {
        this.getUser();
    },
    components: { GvbNav }
}
</script>