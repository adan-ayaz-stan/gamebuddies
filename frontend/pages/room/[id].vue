<template>
  <Sidebar />
  <div class="min-h-screen flex items-stretch md:pl-16">
    <div class="flex-1 p-4 md:p-8 space-y-8 border-x-2 border-lightRed">
      <h2>Room ID: {{ route.params.id }}</h2>

      <div class="relative grid grid-cols-6 gap-4 pb-20">
        <div v-for="i in 10" class="border border-lightRed py-5 px-4 flex flex-col items-center">
          <div class="h-20 w-20 bg-lightRed rounded-full" />
          <h3 class="mt-3">User Name {{ i }}</h3>

          <!-- If user is speaking or not -->
          <div class="w-full h-8 border border-fadeRed relative">

          </div>

          <!-- Dropdown Controls -->
          <UDropdown class="mt-4" :items="items" :popper="{ placement: 'bottom-start' }">
            <UButton color="white" label="Actions" trailing-icon="i-heroicons-chevron-down-20-solid" />
          </UDropdown>
        </div>
      </div>

      <!-- Controls -->
      <div
        class="fixed left-1/2 flex flex-row gap-2 items-center justify-between bottom-10 py-2 rounded px-4 bg-darkGray">
        <!-- Mute -->
        <div class="p-2 bg-white rounded-full text-lightRed cursor-pointer hover:opacity-80">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="M12 14q-1.25 0-2.125-.875T9 11V5q0-1.25.875-2.125T12 2t2.125.875T15 5v6q0 1.25-.875 2.125T12 14m-1 6v-2.075q-2.3-.325-3.937-1.95t-1.988-3.95q-.05-.425.225-.725T6 11t.713.288T7.1 12q.35 1.75 1.738 2.875T12 16q1.8 0 3.175-1.137T16.9 12q.1-.425.388-.712T18 11t.7.3t.225.725q-.35 2.275-1.975 3.925T13 17.925V20q0 .425-.288.713T12 21t-.712-.288T11 20" />
          </svg>
        </div>
        <!-- End Call -->
        <div class="p-2 bg-lightRed rounded-full text-white cursor-pointer hover:opacity-80">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="M12 8q2.95 0 5.813 1.188T22.9 12.75q.3.3.3.7t-.3.7l-2.3 2.25q-.275.275-.638.3t-.662-.2l-2.9-2.2q-.2-.15-.3-.35t-.1-.45v-2.85q-.95-.3-1.95-.475T12 10t-2.05.175T8 10.65v2.85q0 .25-.1.45t-.3.35l-2.9 2.2q-.3.225-.663.2t-.637-.3l-2.3-2.25q-.3-.3-.3-.7t.3-.7q2.2-2.375 5.075-3.562T12 8" />
          </svg>
        </div>

      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const route = useRoute();

const items = [
  [{
    label: 'Profile',
    avatar: {
      src: 'https://avatars.githubusercontent.com/u/739984?v=4'
    }
  }], [{
    label: 'Edit',
    icon: 'i-heroicons-pencil-square-20-solid',
    shortcuts: ['E'],
    click: () => {
      console.log('Edit')
    }
  }, {
    label: 'Duplicate',
    icon: 'i-heroicons-document-duplicate-20-solid',
    shortcuts: ['D'],
    disabled: true
  }], [{
    label: 'Archive',
    icon: 'i-heroicons-archive-box-20-solid'
  }, {
    label: 'Move',
    icon: 'i-heroicons-arrow-right-circle-20-solid'
  }], [{
    label: 'Delete',
    icon: 'i-heroicons-trash-20-solid',
    shortcuts: ['⌘', 'D']
  }]
]

</script>

<style></style>