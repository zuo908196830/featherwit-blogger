<template>
    <div>
        <div>
            <el-input type="textarea" autosize placeholder="请输入题目" v-model="title">
            </el-input>
        </div>
        <mavon-editor v-model="content" ref="md" @change="change" style="min-height: 600px" />
        <button @click="submit">提交</button>
    </div>
</template>

<script>
import axios from 'axios';
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

export default {
    // 注册
    components: {
        mavonEditor,
    },
    data() {
        return {
            content: '', // 输入的markdown
            html: '',    // 及时转的html
            title: ''
        }
    },
    methods: {
        // 所有操作都会被解析重新渲染
        change(value, render) {
            // render 为 markdown 解析后的结果[html]
            this.html = render;
        },
        // 提交
        submit() {
            console.log(this.content);
            console.log(this.html);
            let blogData = {}
            blogData.title = this.title
            blogData.content = this.content
            axios.post('/api/blog/add', blogData).then(res => {
                if (res.data.code === 0) {
                    // todo 成功后跳转文章列表
                    this.$message({
                        message: '恭喜你，这是一条成功消息',
                        type: 'success'
                    });
                    this.$router.push('/')
                }
            }).catch(() => {

            })
        }
    },
    mounted() {
        if (localStorage.getItem('loginStatus') === 'false') {
            this.$message({
                message: '请先登录',
                type: 'warning'
            });
            this.$router.push('/login')
        }
    }
}
</script>