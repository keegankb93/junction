# frozen_string_literal: true

require "bundler/gem_tasks"
require "fileutils"
require_relative "lib/junction"
require_relative "lib/junction/platform"

begin
  require "minitest/test_task"

  Minitest::TestTask.create
rescue LoadError
  warn "Skipping test task because minitest is not installed."
end

begin
  require "rubocop/rake_task"

  RuboCop::RakeTask.new
rescue LoadError
  warn "Skipping rubocop task because rubocop is not installed."
end

namespace :tui do
  desc "Build the Rust TUI binary for the current platform"
  task :build do
    output = Junction::Platform.tui_binary_path(root: __dir__)
    cargo = ENV.fetch("CARGO", "cargo")
    target_dir = File.join(__dir__, ".cache", "rust-target")
    built_binary = File.join(target_dir, "release", "junction-rust-tui")

    FileUtils.mkdir_p(File.dirname(output))

    Dir.chdir(File.join(__dir__, "rust")) do
      sh cargo, "build", "--release", "--target-dir", target_dir
    end

    FileUtils.cp(built_binary, output)
    FileUtils.chmod(0o755, output)
  end
end

default_tasks = []
default_tasks << :test if Rake::Task.task_defined?(:test)
default_tasks << :rubocop if Rake::Task.task_defined?(:rubocop)

task default: default_tasks
