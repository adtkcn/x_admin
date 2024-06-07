declare module 'vue' {
    interface ComponentCustomProperties {
        $filePath: (typeof import('@/methods/index'))['filePath'],
        $perms: (typeof import('@/utils/perms'))['perms']

    }
}

export {}