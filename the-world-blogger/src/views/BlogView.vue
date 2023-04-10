<template>
    <div>
        <div style="position:absolute; left:7%; weight:86%">
            <div>
                <h2 style="text-align: center;">{{ blog.title }}</h2>
            </div>
            <div>
                <VueMarkDown :source="blog.content" />
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import VueMarkDown from 'vue-markdown'

export default {
    components: {
        VueMarkDown
    },
    data() {
        return {
            id: 0,
            blog: {}
        }
    },
    methods: {
        getBlog() {
            axios.get('api/blog/id/' + this.id).then(res => {
                if (res.data.code === 0) {
                    this.blog = res.data.data
                } else if (res.data.code === 1003) {
                    this.$message({
                        message: '文章不存在',
                        type: 'error'
                    })
                }
            }).catch(() => {
                this.$message({
                    message: '服务器错误',
                    type: 'error'
                })
            })
        }
    },
    mounted() {
        this.id = this.$route.query.id
        this.getBlog()
    }
}
</script>