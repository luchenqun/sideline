<template>
  <div class="container" style="padding:5px 15px">
    <div class="detail-box" style="padding;:0px">
      <div class="detail3" style="padding:0px;">
        <el-row :gutter="0">
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Keplr</div>
              <div v-if="address" class="value">{{balance}}</div>
              <div v-else class="value">Please connect keplr</div>
            </div>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="12" :xl="12">
            <div class="data">
              <div class="key">Validator</div>
              <div class="value">{{validatorBalance}}</div>
            </div>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>

  <div class="container">
    <div class="gva-btn-list">
      <el-button type="primary" :disabled="!address" @click="dialogVisible = true;">Regist Developer</el-button>
    </div>
    <!-- <el-divider /> -->
    <el-table :data="developers">
      <el-table-column align="left" min-width="100" label="avatar">
        <template #default="scope">
          <el-avatar :src="scope.row.avatar"> <img src="https://cube.elemecdn.com/e/fd/0fc7d20532fdaf769a25683617711png.png" /> </el-avatar>
        </template>
      </el-table-column>
      <el-table-column align="left" show-overflow-tooltip label="name" min-width="100" prop="name">
        <template #default="scope">
          <div class="nowrap">
            <router-link show-overflow-tooltip :to="('/developer/'+scope.row.address)">{{scope.row.name}}</router-link>
          </div>
        </template>
      </el-table-column>
      <el-table-column align="left" label="balance" show-overflow-tooltip min-width="100" prop="balance" />
      <el-table-column align="left" label="introduce" show-overflow-tooltip min-width="100" prop="introduce" />
      <el-table-column align="left" label="email" show-overflow-tooltip min-width="100" prop="email" />
      <el-table-column align="left" label="address" show-overflow-tooltip min-width="180" prop="address" />
      <el-table-column align="left" label="education" show-overflow-tooltip min-width="100" prop="education" />
      <el-table-column align="left" label="major" show-overflow-tooltip min-width="100" prop="major" />
      <el-table-column align="left" width="80" label="task">
        <template #default="scope">{{ scope.row.taskIds.length }}</template>
      </el-table-column>
      <el-table-column align="left" show-overflow-tooltip min-width="150" label="skills">
        <template #default="scope">
          <el-tag style="margin-right:3px;" v-for="skill of scope.row.skills" :key="skill">{{ skill }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="right" label="action" width="88">
        <template #default="scope">
          <el-button icon="tickets" size="small" type="primary" link @click="toDetailDeveloper(scope.row)">DETAIL</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :before-close="closeDialog" title="Regist Developer">
      <el-alert :title="'It takes ' + registrationFee + ' to regist and give the validator as a reward'" type="warning" effect="dark" style="margin:-20px 0px 20px 0px" />
      <el-form ref="formRef" :model="form" :rules="rules" label-width="90px">
        <el-form-item label="creator" prop="creator">
          <el-input v-model="form.creator" disabled autocomplete="off" />
        </el-form-item>
        <el-form-item label="name" prop="name">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="avatar" prop="avatar">
          <el-input v-model="form.avatar" placeholder="please input you avatar http url" autocomplete="off" />
        </el-form-item>
        <el-form-item label="email" prop="email">
          <el-input v-model="form.email" autocomplete="off" />
        </el-form-item>
        <el-form-item label="education" prop="education">
          <el-input v-model="form.education" autocomplete="off" />
        </el-form-item>
        <el-form-item label="major" prop="major">
          <el-input v-model="form.major" autocomplete="off" />
        </el-form-item>
        <el-form-item label="skills" prop="skills">
          <el-select style="width:100%" v-model="form.skills" multiple filterable allow-create default-first-option :reserve-keyword="false" placeholder="Choose or add skills for you">
            <el-option v-for="item in skillOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="introduce" prop="introduce">
          <v-md-editor v-model="form.introduce" height="400px"></v-md-editor>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="submitRegistDeveloper">OK</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
import { computed, onBeforeMount, ref, watch } from 'vue'
import { useStore } from 'vuex'
import { ElMessage, ElLoading } from 'element-plus';
import { useRouter } from 'vue-router';

export default {
  name: 'Developers',

  setup() {
    // store
    const $s = useStore()
    const router = useRouter();
    const validator = ref("sl1vmuufd2aen7wkeqz7ya26f90l8c62ey9083rax")
    const validatorBalance = ref("")
    const sideline = ref("sl1v7kq6jje3eq43k7eeue7vjjn9lpp0fvrmdw9wm")
    const sidelineBalance = ref("")
    const address = computed(() => $s.getters['common/wallet/address'])
    const balance = ref("")
    const developers = ref([])
    const init = { name: "luke", avatar: "https://b.lucq.fun/images/admin.jpg", email: "luke@qq.com", introduce: "", education: "长沙理工大学", major: "计算机科学与技术", skills: ["React"] }
    const form = ref({})
    const formRef = ref(null);
    const dialogVisible = ref(false)
    const rules = ref({
      creator: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      name: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      introduce: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      email: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      avatar: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      education: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      major: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
      skills: [{ required: true, message: 'This item must be filled in', trigger: 'blur' }],
    });
    const registrationFee = ref("");
    const skillOptions = [
      {
        value: 'JavaScript',
        label: 'JavaScript',
      },
      {
        value: 'Vue',
        label: 'Vue',
      },
      {
        value: 'React',
        label: 'React',
      },
      {
        value: 'Cosmos',
        label: 'Cosmos',
      },
    ]
    const denom = "wrmb";

    onBeforeMount(async () => {
      getData()
      form.value.creator = address.value
    });

    watch(() => address.value, async () => {
      form.value.creator = address.value
      if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }
    })

    const getData = async () => {
      let reply = await $s.dispatch('sideline.sideline/QueryDeveloperAll', {});
      console.log("QueryDeveloperAll", reply)
      for (let developer of reply.developer) {
        const { balance } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: developer.address }, query: { denom } })
        console.log(balance)
        developer.balance = balance.amount + denom;
      }
      developers.value = reply.developer
      reply = await $s.dispatch('sideline.sideline/QueryParams', {});
      // params.value = reply.params
      registrationFee.value = reply?.params?.registrationFee || ""

      if (address.value) {
        const { balance: bal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: address.value }, query: { denom } })
        balance.value = bal.amount + denom;
      }

      const { balance: vbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: validator.value }, query: { denom } })
      validatorBalance.value = vbal.amount + denom;

      const { balance: sbal } = await $s.dispatch('cosmos.bank.v1beta1/QueryBalance', { params: { address: sideline.value }, query: { denom } })
      sidelineBalance.value = sbal.amount + denom;
    }

    const initForm = () => {
      formRef.value.resetFields();
      form.value = { creator: address.value };
    };

    const closeDialog = () => {
      dialogVisible.value = false;
      initForm()
    }

    const submitRegistDeveloper = () => {
      formRef.value.validate(async (valid) => {

        if (valid && form.value.introduce != "") {
          const { value } = form
          console.log("value", value)
          const loading = ElLoading.service({ lock: true, text: 'registing developer...' });
          try {
            const fee = [{ denom: "wrmb", amount: "0" }]
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

    const toDetailDeveloper = (row) => {
      router.push({
        name: 'developer',
        params: row,
      });
    }

    return {
      address,
      balance,
      validator,
      validatorBalance,
      sideline,
      sidelineBalance,
      developers,
      dialogVisible,
      rules,
      form,
      formRef,
      registrationFee,
      skillOptions,
      getData,
      initForm,
      closeDialog,
      submitRegistDeveloper,
      toDetailDeveloper
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/detail.scss";
</style>
