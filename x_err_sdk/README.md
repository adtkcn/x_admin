        <script src="/XErr.umd.js"></script>
        <script>
            const xErr = new XErr.Base(
                {
                    Dns: `${location.origin}/api`,
                    Pid: 'e19e3be20de94f49b68fafb4c30668bc',
                    Uid: ''
                },
                new XErr.Web({})
            )
            xErr.SetUid(2) //设置用户ID
        </script>