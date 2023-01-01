<template>
  <div class="container">
    <div class="detail-box">
      <div class="detail1">
        <div class="left">
          <el-avatar :src="developer.avatar"> <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" /> </el-avatar>
        </div>
        <div class="right">
          <div class="name">{{developer.name}}</div>
          <div>
            <el-row :gutter="0">
              <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                <div class="data">
                  <div class="key">address</div>
                  <div class="value">{{developer.address}}</div>
                </div>
              </el-col>
              <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
                <div class="data">
                  <div class="key">email</div>
                  <div class="value">{{developer.email}}</div>
                </div>
              </el-col>
            </el-row>
          </div>
        </div>
      </div>

      <div class="detail2">
        <el-divider class="detail-divider" />
        <div class="items">
          <div class="item">
            <div class="left">
              <img src="@/assets/project_total.png" alt />
            </div>
            <div class="right">
              <div class="value-special"><span class="c1">{{developer?.taskIds?.length}}</span></div>
              <div class="key">Task Total</div>
            </div>
          </div>
          <el-divider direction="vertical" class="divider" />
          <div class="item">
            <div class="left">
              <img src="@/assets/project_success.png" alt />
            </div>
            <div class="right">
              <div class="value-special"><span class="c2">{{developer?.feedbacks?.length}}</span></div>
              <div class="key">Feedback Total</div>
            </div>
          </div>
        </div>
        <el-divider class="detail-divider" />
      </div>
      <div class="key" style="margin-top:15px;">developer introduce</div>
      <v-md-preview :text="developer.introduce" ref="preview"></v-md-preview>
      <el-divider class="detail-divider" />
    </div>
  </div>

  <div class="container">
    <el-table :data="tasks">
      <el-table-column align="left" show-overflow-tooltip label="title" min-width="100" prop="title">
        <template #default="scope">
          <div class="nowrap">
            <router-link show-overflow-tooltip :to="('/task/'+scope.row.id)">{{scope.row.title}}</router-link>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="description" show-overflow-tooltip min-width="100" prop="description" />
      <el-table-column align="left" label="employer" show-overflow-tooltip min-width="100" prop="employer" />
      <el-table-column align="left" label="remuneration" show-overflow-tooltip min-width="100" prop="remuneration" />
      <el-table-column align="left" label="deposit" show-overflow-tooltip min-width="100" prop="deposit" />
      <el-table-column align="left" label="collateral" show-overflow-tooltip min-width="100" prop="collateral" />
      <el-table-column align="left" label="deadline" show-overflow-tooltip min-width="100" prop="deadline" />
      <el-table-column align="left" min-width="100" label="status">
        <template #default="scope">{{ formatTaskStatus(scope.row.status) }}</template>
      </el-table-column>
      <el-table-column align="right" label="action" width="88">
        <template #default="scope">
          <el-button icon="tickets" size="small" type="primary" link @click="toDetailEmployer(scope.row)">DETAIL</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { computed, onBeforeMount, ref, watch, reactive } from 'vue'
import { useStore } from 'vuex'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage, ElLoading } from 'element-plus';
import { useRoute } from 'vue-router';
import { formatTaskStatus } from '@/utils/sideline';

export default {
  name: 'Developer',

  setup() {
    // store
    const $s = useStore()
    const route = useRoute();
    const { address } = route.params
    const walletAddress = computed(() => $s.getters['common/wallet/address'])
    const tasks = ref([])
    const developer = ref({})
    const preview = ref();

    onBeforeMount(async () => {
      getData()
    });

    watch(() => walletAddress.value, async () => {
      developer.value.creator = address.value
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryDeveloper', { params: { index: address } });
      developer.value = reply.developer
      console.log(reply.developer)

      let items = []
      for (const id of reply.developer.taskIds) {
        let reply = await $s.dispatch('sideline.sideline/QueryTask', { params: { id } });
        items.push(reply.Task)
      }
      tasks.value = items
    }

    const initForm = () => {
      formRef.value.resetFields();
      developer.value = { creator: address.value, introduce: '' };
    };

    const closeDialog = () => {
      initForm()
    }

    const submitRegistDeveloper = () => {
      formRef.value.validate(async (valid) => {

        if (valid && developer.value.introduce != "") {
          const { value } = developer
          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'registing developer...' });
          try {
            const fee = [{ denom: "wrmb", amount: "20000000" }]
            const reply = await $s.dispatch("sideline.sideline/sendMsgRegistDeveloper", { value, fee });
            console.log("reply", reply)
            if (reply.code == 0) {
              ElMessage({
                type: 'success',
                message: 'regist developer success',
              })
              await getData()
              closeDialog()
            } else {
              ElMessage({
                type: 'error',
                message: reply.rawLog,
              })
            }
          } catch (error) {
            console.log("regist developer error", error)
            ElMessage({
              type: 'error',
              message: error,
            })
          }
          loading.close()
        } else {
          ElMessage({
            type: 'warning',
            message: 'please check you input',
          })
        }
      });
    }

    return {
      address,
      tasks,
      preview,
      developer,
      Plus,
      getData,
      initForm,
      closeDialog,
      formatTaskStatus,
      submitRegistDeveloper,
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/detail.scss";
.gva-search-box {
  .el-form-item {
    margin-bottom: 5px;
  }
}
</style>

<style lang="scss">
.container {
  .github-markdown-body {
    padding: 16px 0px 0px 0px;
    font-size: 16px;
    font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
      sans-serif, Apple Color Emoji, Segoe UI Emoji;
    line-height: 1.5;
    word-wrap: break-word;
  }
}
</style>
