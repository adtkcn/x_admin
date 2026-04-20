export type WithIdPid = {
    id?: string | number
    pid?: string | number
}

// ---------- 1. 查询匹配节点 ----------
export function findMatchedNodes(
    data: Record<string, any>[],
    keyword: string,
    fields: string[] = ['name'],
    exact = false
): Record<string, any>[] {
    if (!keyword) return []
    return data.filter((node) =>
        fields.some((field) => {
            const val = node[field]
            return typeof val === 'string' && (exact ? val === keyword : val.includes(keyword))
        })
    )
}

// ---------- 2. 查询所有子级（后代） ----------
export function findDescendants(
    data: Record<string, any>[],
    rootIds: (string | number)[],
    idField = 'id',
    pidField = 'pid'
): Record<string, any>[] {
    // 构建 pid -> children 映射
    const childrenMap = new Map<string | number, Record<string, any>[]>()
    for (const node of data) {
        const p = node[pidField] ?? '__null__'
        if (!childrenMap.has(p)) childrenMap.set(p, [])
        childrenMap.get(p)!.push(node)
    }

    const result: Record<string, any>[] = []
    const queue: (string | number)[] = [...rootIds]
    const visited = new Set<string | number>()

    while (queue.length > 0) {
        const id = queue.shift()!
        if (visited.has(id)) continue
        visited.add(id)

        const children = childrenMap.get(id) || []
        for (const child of children) {
            result.push(child)
            queue.push(child[idField])
        }
    }

    return Array.from(new Map(result.map((n) => [n[idField], n])).values())
}

// ---------- 3. 查询所有父级（祖先） ----------
export function findAncestors(
    data: Record<string, any>[],
    startIds: (string | number)[],
    idField = 'id',
    pidField = 'pid'
): Record<string, any>[] {
    // 构建 id -> node 映射
    const nodeMap = new Map<string | number, Record<string, any>>()
    for (const node of data) {
        nodeMap.set(node[idField], node)
    }

    const result: Record<string, any>[] = []
    const visited = new Set<string | number>()

    for (const id of startIds) {
        let currentId: string | number | null | undefined = id
        while (currentId != null) {
            const node = nodeMap.get(currentId)
            if (!node || visited.has(currentId)) break

            visited.add(currentId)
            // 跳过起始节点自身（只查“父级”）
            if (currentId !== id) {
                result.push(node)
            }

            currentId = node[pidField]
        }
    }

    return Array.from(new Map(result.map((n) => [n[idField], n])).values())
}

// ---------- 4. 合并所有层级 ----------
export function collectAllLevels(
    matched: Record<string, any>[],
    ancestors: Record<string, any>[],
    descendants: Record<string, any>[],
    idField = 'id'
): Record<string, any>[] {
    const all = [...matched, ...ancestors, ...descendants]
    return Array.from(new Map(all.map((n) => [n[idField], n])).values())
}

// ---------- 5. 主入口：一键获取结构化结果 ----------
export function queryHierarchy(
    data: Record<string, any>[],
    keyword: string,
    options: {
        fields?: string[]
        idField?: string
        pidField?: string
        exact?: boolean
    } = {}
) {
    const { fields = ['name'], exact = false, idField = 'id', pidField = 'pid' } = options

    const matched = findMatchedNodes(data, keyword, fields, exact) // 匹配的节点
    const matchedIds = matched.map((n) => n[idField] as string | number)

    const descendants = findDescendants(data, matchedIds, idField, pidField) // 后代
    const ancestors = findAncestors(data, matchedIds, idField, pidField) // 祖先

    const all = collectAllLevels(matched, ancestors, descendants, idField)
    console.log({
        matched, // 匹配的节点
        ancestors, // 所有父级（不包含 matched 自身）
        descendants, // 所有子级
        all // 合并后的完整列表（去重）
    })

    return {
        matched, // 匹配的节点
        ancestors, // 所有父级（不包含 matched 自身）
        descendants, // 所有子级
        all // 合并后的完整列表（去重）
    }
}
