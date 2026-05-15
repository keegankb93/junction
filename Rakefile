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
  desc "Build the Go TUI binary for the current platform"
  task :build do
    output = Junction::Platform.tui_binary_path(root: __dir__)
    go = ENV.fetch("GO", "go")

    FileUtils.mkdir_p(File.dirname(output))
    ENV["GOCACHE"] ||= File.join(__dir__, ".cache", "go-build")

    Dir.chdir(File.join(__dir__, "go")) do
      sh go, "build", "-o", output, "."
    end
  end
end

default_tasks = []
default_tasks << :test if Rake::Task.task_defined?(:test)
default_tasks << :rubocop if Rake::Task.task_defined?(:rubocop)

task default: default_tasks
