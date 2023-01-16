<template>
    <div>
        <el-form :model="registerForm" status-icon :rules="rules" ref="ruleForm" label-width="100px"
            class="demo-ruleForm">
            <el-form-item label="账号" prop="username">
                <el-input type="text" v-model="registerForm.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="registerForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
                <el-input type="password" v-model="registerForm.checkPass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="昵称">
                <el-input type="text" v-model="registerForm.nickname"></el-input>
            </el-form-item>
            <el-form-item label="电话">
                <el-input type="text" v-model="registerForm.telephone"></el-input>
            </el-form-item>
            <el-form-item label="邮箱">
                <el-input type="text" v-model="registerForm.mail"></el-input>
            </el-form-item>
            <el-form-item label="个人简介">
                <el-input type="text" v-model="registerForm.profile"></el-input>
            </el-form-item>
            <div style="position:absolute; left:45%">
                <el-button @click="regist" type="primary">注册</el-button>
            </div>
        </el-form>
    </div>
</template>

<script>
import axios from 'axios';
import LoginFunc from '../js/LoginFunc';

export default {
    data() {
        var checkUsername = (rule, value, callback) => {
            if (!value) {
                return callback(new Error('账号不能为空'));
            }
        };
        var validatePass = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请输入密码'));
            } else {
                if (this.ruleForm.checkPass !== '') {
                    this.$refs.ruleForm.validateField('checkPass');
                }
                callback();
            }
        };
        var validatePass2 = (rule, value, callback) => {
            if (value === '') {
                callback(new Error('请再次输入密码'));
            } else if (value !== this.ruleForm.pass) {
                callback(new Error('两次输入密码不一致!'));
            } else {
                callback();
            }
        };
        return {
            registerForm: {
                pass: '',
                checkPass: '',
                username: '',
                nickname: '',
                mail: '',
                telephone: '',
                profile: ''
            },
            rules: {
                pass: [
                    { required: true, validator: validatePass, trigger: 'blur' }
                ],
                checkPass: [
                    { required: true, validator: validatePass2, trigger: 'blur' }
                ],
                username: [
                    { required: true, validator: checkUsername, trigger: 'blur' }
                ]
            }
        };
    },
    methods: {
        submitForm(formName) {
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    alert('submit!');
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
        resetForm(formName) {
            this.$refs[formName].resetFields();
        },
        warning() {
            this.$message({
                message: '该账号已被注册',
                type: 'warning'
            });
        },
        regist() {
            let registData = {}
            registData.password = this.registerForm.pass
            registData.username = this.registerForm.username
            registData.nickname = this.registerForm.nickname
            registData.telephone = this.registerForm.telephone
            registData.mail = this.registerForm.mail
            registData.profile = this.registerForm.profile
            registData.role = 1
            axios.post("/api/user/register", registData).then(res => {
                if (res.data.code === 0) {
                    if (res.data.data.nickname !== "") {
                        LoginFunc.setLoginStatus(res.data.data.nickname, res.data.data.token)
                    } else {
                        LoginFunc.setLoginStatus(res.data.data.username, res.data.data.token)
                    }
                    this.$router.push("/")
                } else if (res.data.code === 1002) {
                    this.warning()
                }
            }).catch(() => {
                // 网络失败时行为
            })
        }
    }
}
</script>