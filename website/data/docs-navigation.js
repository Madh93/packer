// The root folder for this documentation category is `pages/docs`
//
// - A string refers to the name of a file
// - A "category" value refers to the name of a directory
// - All directories must have an "index.mdx" file to serve as
//   the landing page for the category

export default [
  '--------',
  'terminology',
  {
    category: 'commands',
    content: ['build', 'console', 'fix', 'fmt', 'inspect', 'validate', 'hcl2_upgrade'],
  },
  {
    category: 'templates',
    content: [
    {
      category: "legacy_json_templates",
      content: [
        'builders',
        'communicator',
        'engine',
        'post-processors',
        'provisioners',
        'user-variables',
      ]
    },
    {
    category: 'hcl_templates',
        content: [
          {
            category: 'blocks',
            content: [
              {
                category: 'build',
                content: [
                  'source',
                  'provisioner',
                  'post-processor',
                  'post-processors',
                ],
              },
              'locals',
              'source',
              'variable',
              'packer',
              'data'
            ],
          },
          {
            category: 'functions',
            content: [
              {
                category: 'contextual',
                content: [
                  'aws_secretsmanager',
                  'consul',
                  'env',
                  'vault',
                ],
              },
              {
                category: 'numeric',
                content: [
                  'abs',
                  'ceil',
                  'floor',
                  'log',
                  'max',
                  'min',
                  'parseint',
                  'pow',
                  'signum',
                ],
              },
              {
                category: 'string',
                content: [
                  'chomp',
                  'format',
                  'formatlist',
                  'indent',
                  'join',
                  'lower',
                  'replace',
                  'regex_replace',
                  'split',
                  'strrev',
                  'substr',
                  'title',
                  'trim',
                  'trimprefix',
                  'trimsuffix',
                  'trimspace',
                  'upper',
                ],
              },
              {
                category: 'collection',
                content: [
                  'chunklist',
                  'coalesce',
                  'coalescelist',
                  'compact',
                  'concat',
                  'contains',
                  'distinct',
                  'element',
                  'flatten',
                  'keys',
                  'length',
                  'lookup',
                  'merge',
                  'range',
                  'reverse',
                  'setintersection',
                  'setproduct',
                  'setunion',
                  'slice',
                  'sort',
                  'values',
                  'zipmap',
                ],
              },
              {
                category: 'encoding',
                content: [
                  'base64decode',
                  'base64encode',
                  'csvdecode',
                  'jsondecode',
                  'jsonencode',
                  'urlencode',
                  'yamldecode',
                  'yamlencode',
                ],
              },
              {
                category: 'file',
                content: [
                  'abspath',
                  'basename',
                  'dirname',
                  'file',
                  'fileexists',
                  'fileset',
                  'pathexpand',
                ],
              },
              {
                category: 'datetime',
                content: ['formatdate', 'timeadd', 'timestamp'],
              },
              {
                category: 'crypto',
                content: [
                  'bcrypt',
                  'md5',
                  'rsadecrypt',
                  'sha1',
                  'sha256',
                  'sha512',
                ],
              },
              {
                category: 'uuid',
                content: ['uuidv4', 'uuidv5'],
              },
              {
                category: 'ipnet',
                content: ['cidrhost', 'cidrnetmask', 'cidrsubnet'],
              },
              {
                category: 'conversion',
                content: ['can', 'convert', 'try'],
              },
            ],
          },
          'variables',
          'locals',
          'contextual-variables',
          'datasources',
          'path-variables',
          'syntax',
          'onlyexcept',
          'expressions',
          'syntax-json',
        ],
      },
    ],
  },
  '----------',
  { category: 'communicators', content: ['ssh', 'winrm'] },
  {
    category: 'builders',
    content: [
      'alicloud-ecs',
      {
        category: 'amazon',
        content: ['chroot', 'ebs', 'ebssurrogate', 'ebsvolume', 'instance'],
      },
      {
        category: 'azure',
        content: ['arm', 'chroot'],
      },
      'cloudstack',
      'digitalocean',
      'docker',
      'file',
      'googlecompute',
      'hetzner-cloud',
      'hyperone',
      { category: 'hyperv', content: ['iso', 'vmcx'] },
      'linode',
      'lxc',
      'lxd',
      'ncloud',
      'null',
      'oneandone',
      'openstack',
      { category: 'oracle', content: ['classic', 'oci'] },
      {
        category: 'outscale',
        content: ['chroot', 'bsu', 'bsusurrogate', 'bsuvolume'],
      },
      { category: 'parallels', content: ['iso', 'pvm'] },
      'profitbricks',
      { category: 'proxmox', content: ['iso', 'clone'] },
      'qemu',
      'scaleway',
      'tencentcloud-cvm',
      'jdcloud',
      'triton',
      'ucloud-uhost',
      'vagrant',
      {
        category: 'virtualbox',
        content: ['iso', 'ovf', 'vm'],
      },
      {
        category: 'vmware',
        content: ['iso', 'vmx', 'vsphere-iso', 'vsphere-clone'],
      },
      'yandex',
      'custom',
      'community-supported',
    ],
  },
  { category: 'datasources', content: ['amazon-ami'] },
  {
    category: 'provisioners',
    content: [
      'ansible-local',
      'ansible',
      'breakpoint',
      'chef-client',
      'chef-solo',
      'converge',
      'file',
      'inspec',
      'powershell',
      'puppet-masterless',
      'puppet-server',
      'salt-masterless',
      'shell',
      'shell-local',
      'windows-shell',
      'windows-restart',
      'custom',
      'community-supported',
    ],
  },
  {
    category: 'post-processors',
    content: [
      'alicloud-import',
      'amazon-import',
      'artifice',
      'compress',
      'checksum',
      'digitalocean-import',
      'docker-import',
      'docker-push',
      'docker-save',
      'docker-tag',
      'exoscale-import',
      'googlecompute-export',
      'googlecompute-import',
      'manifest',
      'shell-local',
      'ucloud-import',
      'vagrant',
      'vagrant-cloud',
      'vsphere',
      'vsphere-template',
      'yandex-export',
      'yandex-import',
      'community-supported',
    ],
  },
  '----------',
  'install',
  '----------',
  {
    category: 'extending',
    content: [
      'plugins',
      'custom-builders',
      'custom-post-processors',
      'custom-provisioners',
      'custom-datasources',
    ],
  },
  '---------',
  'environment-variables',
  'core-configuration',
  'debugging',
]
