# frozen_string_literal: true

require_relative "lib/junction/version"

Gem::Specification.new do |spec|
  spec.name = "junction"
  spec.version = Junction::VERSION
  spec.authors = ["Keegankb93"]
  spec.email = ["keegankb@gmail.com"]

  spec.summary = "A terminal dev hub for Rails applications."
  spec.description = "Junction provides Rails-aware commands and launches a bundled TUI for local development workflows."
  spec.homepage = "https://github.com/keegankb93/junction"
  spec.license = "MIT"
  spec.required_ruby_version = ">= 3.2.0"

  spec.metadata["homepage_uri"] = spec.homepage
  spec.metadata["source_code_uri"] = spec.homepage
  spec.metadata["changelog_uri"] = "#{spec.homepage}/blob/main/CHANGELOG.md"

  # Specify which files should be added to the gem when it is released.
  # The `git ls-files -z` loads the files in the RubyGem that have been added into git.
  spec.files = Dir.chdir(__dir__) do
    Dir[
      "CHANGELOG.md",
      "CODE_OF_CONDUCT.md",
      "LICENSE.txt",
      "README.md",
      "exe/*",
      "lib/**/*",
      "sig/**/*",
      "vendor/bin/*"
    ]
  end
  spec.bindir = "exe"
  spec.executables = spec.files.grep(%r{\Aexe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]

  # Uncomment to register a new dependency of your gem
  # spec.add_dependency "example-gem", "~> 1.0"

  # For more information and examples about making a new gem, check out our
  # guide at: https://bundler.io/guides/creating_gem.html
end
