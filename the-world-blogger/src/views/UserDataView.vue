<template>
    <div>
        <el-descriptions class="margin-top" title="带边框列表" :column="3" :size="size" border>
            <template slot="extra">
                <el-button type="primary" size="small">操作</el-button>
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
</template>

<script>
import axios from 'axios'
export default {
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
        }
    },
    methods: {
        getUser() {
            axios.get("/api/user/data").then(res => {
                if (res.data.code === 0) {
                    this.userData = res.data.data
                } else if (res.data.code === 1006) {
                    // todo 提示未登录
                    this.$message({
                        message: '您还未登录',
                        type: 'warning'
                    })
                    this.$router.push('/login')
                }
            }).catch(() => {

            })
        }
    },
    created() {
        this.getUser()
    }
}
</script>