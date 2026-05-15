# frozen_string_literal: true

require "rbconfig"

module Junction
  class Error < StandardError; end unless const_defined?(:Error)

  module Platform
    module_function

    def tui_binary_path(root:)
      File.join(root, "vendor", "bin", tui_binary_name)
    end

    def tui_binary_name
      "junction-tui-#{tui_platform}"
    end

    def tui_platform
      "#{os}-#{arch}"
    end

    def os
      host_os = RbConfig::CONFIG.fetch("host_os")

      case host_os
      when /darwin/
        "darwin"
      when /linux/
        "linux"
      else
        raise Error, "Unsupported OS: #{host_os}"
      end
    end

    def arch
      host_cpu = RbConfig::CONFIG.fetch("host_cpu")

      case host_cpu
      when /arm64|aarch64/
        "arm64"
      when /x86_64|amd64/
        "amd64"
      else
        raise Error, "Unsupported CPU architecture: #{host_cpu}"
      end
    end
  end
end
