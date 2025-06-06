export default {
    es: {
        keywordPlaceholder: 'host / name / code',
        port: 'Port',
        size: 'size',
        docs: 'docs',
        health: 'health',
        aliases: 'Aliases',
        addAlias: 'Add Alias',
        specifyIdAdd: 'Specify the ID added, if id exists, then update',
        addIndex: 'Add Index',
        editIndex: 'Edit Index',
        status: 'status',
        acName: 'Credential',
        emptyTable: 'data not fund',
        connSuccess: 'be connected successfully',
        shouldTestConn: 'please test connection first',
        instance: 'ES Instance',
        instanceSave: 'Save Instance',
        instanceDel: 'Delete Instance',
        operation: 'Data Operation',
        dataSave: 'Data Save',
        dataDel: 'Data Del',
        indexName: 'Index Name',
        requireIndexName: 'Index Name Is Required',
        indexDetail: 'Index Detail',
        indexMapping: 'Mappings',
        indexStats: 'Stats',
        opViewColumns: 'Option View Columns',
        opIndex: 'Index Management',
        opSearch: 'Search',
        searchParamsPreview: 'Search Params Preview',
        opBasicSearch: 'Basic Search',
        opSeniorSearch: 'Senior Search',
        sampleMappings: 'Sample Mappings',
        copyMappings: 'Copy Mappings',
        readonlyMsg: 'The content is readOnly',
        opDashboard: 'Dashboard',
        opSettings: 'Settings',
        templates: 'Templates',
        availableSettingFields: 'Available Setting Fields',
        Reindex: 'Reindex',
        ReindexTargetIdx: 'Target Index',
        ReindexIsSync: 'Sync Able',
        ReindexDescription:
            "If a field in Mapping has been defined, you can't modify the type of the field, and you can't change the number of shards, you can use the Reindex API to solve this problem.",
        ReindexSyncDescription: 'If the amount of index data is large, we recommend that you enable asynchronous data to avoid request timeouts.',
        ReindexToOtherInst: 'To other Instance',
        ReindexSyncTask: 'Sync Task',
        makeSearchParam: 'Make Search Params',
        filterColumn: 'Filter Columns',
        searchParams: 'Search',
        searchParamsDesc: 'If no field is selected or no condition value is set, it will not take effect',
        standardSearch: 'Standard Search',
        AggregationSearch: 'Aggregation Search',
        SqlSearch: 'Sql Search',
        searchError: 'Search Error',
        execError: 'Exec Error',
        docJsonError: 'Document JSON Format Error',
        sortParams: 'Sort',
        otherParams: 'Other',
        closeIndexConfirm: 'This operation will close index [{name}]. Do you want to continue?',
        openIndexConfirm: 'This operation will open index [{name}]. Do you want to continue?',
        clearCacheConfirm: 'This operation will clear index [{name}] cache. Do you want to continue?',
        page: {
            home: 'First Page',
            prev: 'Prev Page',
            next: 'Next Page',
            total: 'Total Count',
            changeSize: 'Change Page Size',
        },
        temp: {
            addTemp: 'Add template',
            view: 'Template Detail',
            name: 'name',
            priority: 'priority',
            index_patterns: 'patterns',
            content: 'content',
            showHide: 'show system templates',
            description: 'description',
            filter: 'filter name / description',
            versionAlert: 'Versions prior to 7.8 are not supported',
            note: `1、When creating a new index, if the index name matches the wildcard of the index template, the index template's settings (_setting, _mapping, etc.) are used。
2、Templates take effect only when an index is created, and modifying a template does not affect existing indexes。
3、You can specify the value of "priority", which was "order" before version 7.8, and if the new index name matches multiple templates, the one with the lowest priority will be used first.`,
        },
        dashboard: {
            instInfo: 'Instance Info',
            clusterHealth: 'Cluster Health',
            nodes: 'Nodes Info',
            sysMem: 'System Mem',
            jvmMem: 'JVM Mem',
            fileSystem: 'File System',
            analyze: 'Analyze',
            idxName: 'Index Name',
            field: 'Field',
            text: 'Text',
            startAnalyze: 'Start Analyze',
        },
        contextmenu: {
            index: {
                addIndex: 'Add Index',
                showSys: 'Show System Index',
                copyName: 'Copy Name',
                refresh: 'Refresh Index',
                flush: 'Flush Index',
                clearCache: 'Clear Index Cache',
                addAlias: 'Add Alias',
                Close: 'Close',
                Open: 'Open',
                Delete: 'Delete Index',
                edit: 'Edit Index',
                DeleteSelectLine: 'Copy Selected Line Json',
                BaseSearch: 'Base Search',
                SeniorSearch: 'Senior Search',
                copyLineJson: 'Copy Line Json',
                copySelectLineJson: 'Copy Selected Line Json',
            },
        },
    },
};
