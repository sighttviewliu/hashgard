module.exports = {
  base: '/github-pages/',
  markdown: {
    lineNumbers: false // 代码块显示行号
  },
  locales: {
    // 键名是该语言所属的子路径
    // 作为特例，默认语言可以使用 '/' 作为其路径。
    '/': {
      title: "Hashgard Documentation",
      lang: 'en-US',
      description: 'Welcome to the hashgard documentation'
    },
    '/translations/zh/': {
      title: "Hashgard Documentation",
      lang: 'zh-CN',
      description: '欢迎来到hashgard文档中心'
    }
  },
  head: [
    ['link', {
      rel: 'icon',
      href: '/favicon.ico'
    }]
  ],
  themeConfig: {
    repo: 'hashgard/github-pages',
    docsRepo: 'hashgard/hashgard',
    docsDir: 'docs',
    editLinks: true,
    // 默认为 "Edit this page"
    editLinkText: '帮助我们改善此页面！',
    locales: {
      '/': {
        selectText: 'Languages',
        label: 'English',
        editLinkText: 'Edit this page on GitHub',
        serviceWorker: {
          updatePopup: {
            message: "New content is available.",
            buttonText: "Refresh"
          }
        },
        algolia: {},
        nav: [{
            text: 'Guide',
            link: '/learn/introduction.md'
          },
          {
            text: 'command',
            link: '/cli/'
          },

        ],
        sidebar: {
          '/learn/': [{
              title: 'Overview',
              children: [
                ['/learn/introduction.md', 'who is Hashgard'],
                ['/learn/howToContribute.md', 'how To Contribute']
              ]
            },
            {
              title: 'UsersGuide',
              children: [
                ['/learn/UsersGuide/WebWalletGuide.md', 'WebWalletGuide'],
                ['/learn/UsersGuide/hashgardExplorerGuide.md', 'ExplorerGuide'],
                ['/learn/UsersGuide/hashgardNav.md', 'hashgardNav']
              ]
            },
            {
              title: 'Guide',
              children: [
                ['/learn/installation.md', 'installation'],
                ['/feature/AccountGuide.md', 'AccountGuide'],
                ['/learn/Guide/', 'testnet Guide'],
              ]
            }
          ],
          '/cli/': [
            ['/cli/', 'cli'],
            {
              title: 'hashgardcli',
              children: [
                ['/cli/hashgardcli/', 'directory'],
                ['/cli/hashgardcli/bank/', 'bank'],
                ['/cli/hashgardcli/distribution/', 'distribution'],
                ['/cli/hashgardcli/gov/', 'gov'],
                ['/cli/hashgardcli/keys/', 'keys'],
                ['/cli/hashgardcli/stake/', 'stake'],
                ['/cli/hashgardcli/status', 'status'],
                ['/cli/hashgardcli/tendermint/', 'tendermint'],
                ['/cli/hashgardcli/slashing/', 'slashing'],
                ['/cli/hashgardcli/issue/', 'issue'],
                ['/cli/hashgardcli/box/', 'box'],
                ['/cli/hashgardcli/faucet/send.md', 'faucet'],
                ['/cli/hashgardcli/exchange/', 'exchange'],
              ]
            },
            {
              title: 'hashgard',
              children: [
                ['/cli/hashgard/', 'directory'],
                ['/cli/hashgard/init', 'hashgard init'],
                ['/cli/hashgard/add-genesis-account', 'hashgard add-genesis-account'],
                ['/cli/hashgard/gentx', 'hashgard gentx'],
                ['/cli/hashgard/collect-gentxs', 'hashgard collect-gentxs'],
                ['/cli/hashgard/validate-genesis', 'hashgard validate-genesis'],
                ['/cli/hashgard/start', 'hashgard start'],
                ['/cli/hashgard/testnet', 'hashgard testnet'],
                ['/cli/hashgard/unsafe-reset-all', 'hashgard unsafe-reset-all'],
                ['/cli/hashgard/export', 'hashgard export'],
                ['/cli/hashgard/tendermint', 'hashgard tendermint'],

              ],
            },
            ['/cli/hashgardlcd/', 'hashgardlcd'],
          ],

          sidebarDepth: 2, // e'b将同时提取markdown中h2 和 h3 标题，显示在侧边栏上。
          lastUpdated: 'Last Updated', // 文档更新时间：每个文件git最后提交的时间,

        },

      },

      '/translations/zh/': {
        selectText: '选择语言',
        label: '简体中文',
        editLinkText: '在 GitHub 上编辑此页',
        serviceWorker: {
          updatePopup: {
            message: "发现新内容可用.",
            buttonText: "刷新"
          }
        },
        // 当前 locale 的 algolia docsearch 选项
        nav: [{
            text: '教程',
            link: '/translations/zh/learn/introduction.md'
          },
          {
            text: '命令',
            link: '/translations/zh/cli/'
          },
        ],
        sidebar: {
          '/translations/zh/cli/': [
            ['/translations/zh/cli/', '命令手册'],
            {
              title: 'hashgardcli',
              children: [
                ['/translations/zh/cli/hashgardcli/', 'directory'],
                ['/translations/zh/cli/hashgardcli/distribution/', 'distribution'],
                ['/translations/zh/cli/hashgardcli/gov/', 'gov'],
                ['/translations/zh/cli/hashgardcli/keys/', 'keys'],
                ['/translations/zh/cli/hashgardcli/stake/', 'stake'],
                ['/translations/zh/cli/hashgardcli/status', 'status'],
                ['/translations/zh/cli/hashgardcli/tendermint/', 'tendermint'],
                ['/translations/zh/cli/hashgardcli/slashing/', 'slashing'],
                ['/translations/zh/cli/hashgardcli/issue/', 'issue'],
                ['/translations/zh/cli/hashgardcli/box/', 'box'],
                ['/translations/zh/cli/hashgardcli/faucet/send.md', 'faucet'],
                ['/translations/zh/cli/hashgardcli/exchange/', 'exchange']

              ],
            },
            {
              title: 'hashgard',
              children: [
                ['/translations/zh/cli/hashgard/', 'directory'],
                ['/translations/zh/cli/hashgard/init.md', 'hashgard init'],
                ['/translations/zh/cli/hashgard/gentx.md', 'hashgard gentx'],
                ['/translations/zh/cli/hashgard/collect-gentxs.md', 'hashgard collect-gentxs'],
                ['/translations/zh/cli/hashgard/validate-genesis.md', 'hashgard validate-genesis'],
                ['/translations/zh/cli/hashgard/start.md', 'hashgard start'],
                ['/translations/zh/cli/hashgard/testnet.md', 'hashgard testnet'],
                ['/translations/zh/cli/hashgard/unsafe-reset-all.md', 'hashgard unsafe-reset-all'],
                ['/translations/zh/cli/hashgard/export.md', 'hashgard export'],
                ['/translations/zh/cli/hashgard/tendermint.md', 'hashgard tendermint'],

              ],
            },
            ['/translations/zh/cli/hashgardlcd/', 'hashgardlcd'],
          ],
          '/translations/zh/learn/': [{
              title: '总览',
              children: [
                ['/translations/zh/learn/introduction.md', 'Hashgard是什么'],
                ['/translations/zh/learn/howToContribute.md', '怎样参与建设'],
              ],
            },
            {
              title: '用户使用指南',
              children: [
                ['/translations/zh/learn/UsersGuide/WebWalletGuide.md', '网页钱包使用指南'],
                ['/translations/zh/learn/UsersGuide/hashgardExplorerGuide.md', '浏览器使用指南'],
                ['/translations/zh/learn/UsersGuide/hashgardNav.md', 'hashgard导航'],

              ],
            },
            {
              title: '教程',
              children: [
                ['/translations/zh/learn/installation.md', '安装hashgard'],
                ['/translations/zh/features/AccountGuide.md', '账户类型说明'],
                ['/translations/zh/learn/Guide/', '测试网络指南'],
              ],
            },
          ],
        },
        sidebarDepth: 2, // e'b将同时提取markdown中h2 和 h3 标题，显示在侧边栏上。
        lastUpdated: 'Last Updated', // 文档更新时间：每个文件git最后提交的时间,

      },

    }
  }
}
