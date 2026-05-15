# frozen_string_literal: true

require "test_helper"

class TestJunction < Minitest::Test
  def test_that_it_has_a_version_number
    refute_nil ::Junction::VERSION
  end

  def test_tui_binary_name_uses_junction_prefix
    require "junction/platform"

    assert_match(/\Ajunction-tui-/, Junction::Platform.tui_binary_name)
  end
end
